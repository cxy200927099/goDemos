package main

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
	"github.com/gocolly/colly"
	"log"
	"strconv"
	"strings"
	"time"
)

const (
	KEY_DIR_KEY string = "dirKey"
	KEY_DIR_NAME string = "dirName"
	KEY_SECOND_KEY string = "secondKey"
)

type JobItem struct {
	DirKey string
	DirName string
	PicUrl	string
}

//save the stars key and name
var starsNameMap = make(map[string]string)
var starsCountMap = make(map[string]int64)

func convert2Enoder(input string, encoder string) string{
	enc := mahonia.NewEncoder(encoder)
	return enc.ConvertString(input)
}

func getKeyFromUrl(url string, label string) (string, error){
	strs := strings.Split(url, label)
	if len(strs) != 2{
		return "", errors.New(fmt.Sprintf("getKeyFromUrl failed! url:%s label:%s\n", url, label))
	}
	subStrs := strings.Split(strs[1], "/")
	keyName := subStrs[1]

	return keyName, nil
}

func getSecondKeyFromUrl(url string, label string) (string, error){
	strs := strings.Split(url, label)
	if len(strs) != 2{
		return "", errors.New(fmt.Sprintf("getSecondKeyFromUrl failed! url:%s label:%s\n", url, label))
	}

	return strs[1], nil
}

func parsePicListOfStars(url string){

}

func main() {
	// Instantiate default collector
	c := colly.NewCollector()
	//不设置中文网页抓取的text是乱码
	c.DetectCharset = true

	// On every a element which has href attribute call callback
	c.OnHTML("#tablen63>tbody>tr", func(e *colly.HTMLElement) {

		//log.Printf("tablen63: -> %s\n", e.Name)
		e.ForEach("td", func(i int, td *colly.HTMLElement) {
			td.DOM.Find("a").Each(func(i int, a *goquery.Selection) {
				attr, exists := a.Attr( "href")
				if exists{
					name := a.Text()
					name = convert2Enoder(name, "utf-8")
					absUrl := e.Request.AbsoluteURL(attr)
					keyName, _ := getKeyFromUrl(absUrl, "n_china")
					secondKeyName, _ := getSecondKeyFromUrl(absUrl, "n_china")
					//log.Printf("td->a: %d ->%s %s %s\n", i, absUrl, convert2Enoder(name, "utf-8"), keyName)
					//访问第二页
					e.Request.Ctx.Put(KEY_DIR_KEY, keyName)
					e.Request.Ctx.Put(KEY_DIR_NAME, name)
					e.Request.Ctx.Put(KEY_SECOND_KEY, secondKeyName)

					ctx := colly.NewContext()
					ctx.Put(KEY_DIR_KEY, keyName)
					ctx.Put(KEY_DIR_NAME, name)
					ctx.Put(KEY_SECOND_KEY, secondKeyName)
					_, exist := starsCountMap[keyName]
					if !exist{
						starsCountMap[keyName] = 0
					}
					if keyName == "2R" {
						log.Printf("page1:%s name:%s keyName:%s secondKey:%s\n", absUrl, name, keyName, secondKeyName)
						//e.Request.Visit(absUrl)
						c.Request("GET", absUrl, nil, ctx, nil)
						//c.Visit(absUrl)
					}
				}else{
					log.Printf("page1: %d -> empty a\n", i)
				}

			})

		})
	})

	// parse page attribute and request more page
	c.OnHTML("div#page_info", func(element *colly.HTMLElement) {
		curPage := int64(0)
		maxPage := int64(0)

		url := element.Request.URL.String()
		secondKey := element.Request.Ctx.Get(KEY_SECOND_KEY)
		if secondKey == ""{
			log.Fatalf("get SecondKey From context failed! url:%s\n", url)
		}

		element.DOM.Find("span[class=current]").Each(func(i int, selection *goquery.Selection) {
			pageStr := selection.Text()
			page, err := strconv.ParseInt(pageStr, 10, 64)
			if err == nil{
				curPage = page
			}
		})

		element.DOM.Find("a").Each(func(i int, selection *goquery.Selection) {
			tmpPageStr := selection.Text()
			tmpPage, err := strconv.ParseInt(tmpPageStr, 10, 64)
			if err != nil{
				tmpPage = 0
			}
			if maxPage < tmpPage{
				maxPage = tmpPage
			}
		})
		log.Printf("parse page2 %s cur:%d total:%d\n", secondKey, curPage, maxPage)
		if curPage == 1{
			r := element.Request
			//r.Ctx.Put("curPage", curPage)
			//r.Ctx.Put("maxPage", maxPage)
			r.Ctx.Put(KEY_SECOND_KEY, secondKey)

			//request the left page
			for i:= curPage+1; i <= maxPage; i++{
				reqUrl := fmt.Sprintf("%s%s%d",r.URL.String(),"&page=",i)
				log.Printf("request %s page:%d url:%s\n", secondKey, i, reqUrl)
				r.Visit(reqUrl)
			}
		}

	})

	// parse second picture list page
	c.OnHTML("div#container", func(e *colly.HTMLElement) {
		parentUrl := e.Request.URL.String()

		dirKey := e.Request.Ctx.Get(KEY_DIR_KEY)
		dirName := e.Request.Ctx.Get(KEY_DIR_NAME)
		if dirKey == "" || dirName == ""{
			log.Fatalf("page2 do not need to parse page:%s", parentUrl)
			return
		}
		secondKey := e.Request.Ctx.Get(KEY_SECOND_KEY)
		if secondKey == ""{
			log.Fatalf("page2 get SecondKey From context failed! url:%s\n", parentUrl)
		}

		e.DOM.Find("div[class=imgholder]").Each(func(i int, selection *goquery.Selection) {
			a := selection.Children()
			if !a.Is("a"){
				return
			}
			postUrl, exist := a.Attr("href")
			if !exist{
				log.Printf("cannot find a[href] attribute\n")
				return
			}
			strs := strings.Split(postUrl, "/")
			if strs[len(strs)-1] == "desktop"{
				log.Printf("page2 except url:%s\n", postUrl)
				return
			}

			absPostUrl := e.Request.AbsoluteURL(postUrl)


			log.Printf("page2 %s postUrl:%s\n", parentUrl, absPostUrl)
			e.Request.Ctx.Put(KEY_DIR_KEY, dirKey)
			e.Request.Ctx.Put(KEY_DIR_NAME, dirName)
			e.Request.Ctx.Put(KEY_SECOND_KEY, secondKey)
			e.Request.Visit(absPostUrl)
		})
	})

	// parse third star poster page,get final poster url
	c.OnHTML("div#image_show", func(e *colly.HTMLElement) {
		parentUrl := e.Request.URL.String()
		dirKey := e.Request.Ctx.Get(KEY_DIR_KEY)
		dirName := e.Request.Ctx.Get(KEY_DIR_NAME)
		if dirKey == "" || dirName == ""{
			log.Fatalf("Error third page:%s", parentUrl)
			return
		}
		secondKey := e.Request.Ctx.Get(KEY_SECOND_KEY)
		if secondKey == ""{
			log.Fatalf("get SecondKey From context failed! url:%s\n", parentUrl)
		}
		e.DOM.Find("img").Each(func(i int, selection *goquery.Selection) {
			tmpUrl, _ := selection.Attr("src")
			posterUrl := e.Request.AbsoluteURL(tmpUrl)

			curCnt, exist := starsCountMap[dirKey]
			if !exist{
				starsCountMap[dirKey] = 0
			}else {
				starsCountMap[dirKey] = curCnt+1
			}

			log.Printf("page3 poster total:%d dirName:%s secondKey:%s url:%s\n",curCnt, dirName, secondKey, posterUrl)
		})
	})

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		RandomDelay: 1 * time.Second,
	})


	c.OnResponse(func(r *colly.Response) {
		//doc, err := goquery.NewDocumentFromReader(bytes.NewReader(r.Body))
		//if err != nil{
		//	log.Fatal(err)
		//}
		//
		//doc.Find("div#page_info").Each(func(i int, selection *goquery.Selection) {
		//	attrID, _ :=  selection.Attr("id")
		//	attrClass, _ := selection.Attr("class")
		//	log.Printf("find %d div id:%s class=%s\n", i, attrID, attrClass)
		//
		//})
	})

	// On every a element which has href attribute call callback
	//c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	//	link := e.Attr("href")
	//	// Print link
	//	log.Printf("Link found: %q -> %s\n", e.Text, link)
	//	// Visit link found on page
	//	// Only those links are visited which are matched by  any of the URLFilter regexps
	//	//c.Visit(e.Request.AbsoluteURL(link))
	//})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		//log.Println("Visiting", r.URL.String())
	})

	// Start scraping on http://www.n63.com/photodir/china.htm
	// 第一页明星列表页面
	c.Visit("http://www.n63.com/photodir/china.htm")
}