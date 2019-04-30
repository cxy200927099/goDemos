package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"goDemos/utils"
	"io"
	"os"
)

var cid_path = flag.String("cid_path", "cid.txt", "the cid path,default is cid.txt")
var cid_unexist_path = flag.String("cid_unexist_path", "cid_unexist.txt", "output cid unexist path,default is cid_unexist.txt")

var cid_map = make(map[string]uint64)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readLine(f *os.File)([]string , error){

	r := bufio.NewReader(f)
	lineCnt := 0
	var lines []string
	var err error
	for {
		var buffer bytes.Buffer

		var l []byte
		var isPrefix bool

		for {
			l, isPrefix, err = r.ReadLine()
			buffer.Write(l)

			// If we've reached the end of the line, stop reading.
			if !isPrefix {
				break
			}

			// If we're just at the EOF, break
			if err != nil {
				break
			}
		}

		if err == io.EOF {
			break
		}

		line := buffer.String()
		lineCnt++
		lines = append(lines, line)
	}

	if err != io.EOF {
		fmt.Printf("readLine > Failed!: %v\n", err)
		return nil, nil
	}
	return lines,nil
}

func readLine1(f *os.File)([]string){

	scanner := bufio.NewScanner(f)
	var lineCnt = 0
	var lines []string
	for scanner.Scan()  {
		lineCnt++
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

func testReadLine(){
	f, err := os.Open(*cid_path)
	check(err)
	defer f.Close()

	fmt.Println("=========reader.ReadLine")
	//lines, err := readLine(f)
	//check(err)
	//for _, line := range lines{
	//	fmt.Println(line)
	//}

	fmt.Println("==========Scanner.ReadLine")
	lines := readLine1(f)
	for _, line := range lines{
		fmt.Println(line)
	}
}

func writeLineToFile(){

}

func get_cid_map(){
	f, err := os.Open(*cid_path)
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lineCnt = 0
	for scanner.Scan()  {
		lineCnt++
		line := scanner.Text()

		if _, ok := cid_map[line]; ok{
			fmt.Println("repeat:",line)
		}else{

			value, ret := utils.Uint64(line)
			if !ret {
				fmt.Println("parse uint64 error!")
				continue
			}
			cid_map[line] = value
		}
	}
	fmt.Println("lineCnt:",lineCnt)
	fmt.Println("cid_map.count:", len(cid_map))

	//for k, v := range cid_map{
	//	fmt.Println("",k,":",v)
	//}

}

func find_unexist_cid(){

	get_cid_map()


	f, err := os.OpenFile(*cid_unexist_path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	check(err)
	defer f.Close()

	for i:= 0; i <= 170960; i++{
		key := fmt.Sprintf("%d",i)
		if _, ok := cid_map[key]; ok{
			continue
		}

		wData := fmt.Sprintf("%d\n",i)
		len_W := len(wData)
		n, err := f.Write([]byte(wData))
		if err == nil && n < len_W {
			err = io.ErrShortWrite
			check(err)
		}
	}

}

func main(){
	flag.Parse()
	fmt.Println("the cid path:",*cid_path)

	find_unexist_cid()
}
