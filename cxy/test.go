package main

import (
	"fmt"
	_ "fmt"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const(
	ADD_TEST = "add"
	QUERY_TEST = "query"
)

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func String(val interface{}) string {
	if val == nil {
		return ""
	}
	re_value := reflect.ValueOf(val)
	for re_value.Kind() == reflect.Ptr {
		re_value = re_value.Elem()
		if !re_value.IsValid() {
			return ""
		}
		val = re_value.Interface()
		if val == nil {
			return ""
		}
		re_value = reflect.ValueOf(val)
	}
	if val == nil {
		return ""
	}

	switch v := val.(type) {
	case string:
		return v
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	default:
		return fmt.Sprint(v)
	}
}

func getCurrentTime()(timeStr string){
	now := time.Now()
	year,month,day := now.Date()
	h,m,s := now.Clock()
	nano := now.Nanosecond()
	fmt.Println("curTime string:",time.Now().String())
	str := fmt.Sprintf("%d-%d-%d-%d:%d:%d.%d",year,month,day,h,m,s,nano)
	//fmt.Println("curTime string1:",str)

	return str
}

func testGetTime(){
	//timeNano := 30 * 24 * 60 * 60 * 1000000000
	timeNano :=  10 * 60 * 1000000000
	fmt.Println("10 min timeNano unixNanao:",timeNano)
	postTime := time.Now().UnixNano()
	fmt.Println("curTime unixNanao:",postTime)
	fmt.Println("curTime unixNanao:",uint64(postTime))
	fmt.Println("10 min before curTime unixNanao:",(uint64(postTime)-uint64(timeNano)) )
	now := time.Now()
	year,month,day := now.Date()
	h,m,s := now.Clock()
	nano := now.Nanosecond()
	fmt.Println("curTime nanoSecond:",nano)
	fmt.Println("curTime string:",time.Now().String())
	str := fmt.Sprintf("%d-%d-%d-%d:%d:%d.%d",year,month,day,h,m,s,nano)
	fmt.Println("curTime string1:",str)

	str2 := now.Format("2018-03-05-18:03:07")
	fmt.Println("curTime string2:",str2)

}

func Substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func TestString(){
	fmt.Println("---------- TestString ---------------")
	videoId := "fda36c63b5d1851d6e4a3bae22dccbad"
	lenStr := len(videoId)
	path1 := Substr(videoId, 0, 2)
	path2 := Substr(videoId, lenStr-2, 2)
	fmt.Println("-----videoId=",videoId)
	fmt.Println("-----path1=",path1," path2=",path2)
}

func TestGetEnv(){
	logLevel := os.Getenv("GRPC_GO_LOG_SEVERITY_LEVEL")
	vLevel := os.Getenv("GRPC_GO_LOG_VERBOSITY_LEVEL")
	fmt.Println("logLevel=",logLevel)
	fmt.Println("vLevel=",vLevel)
}

func TestCalcute(){

	//总量特征 =10 * 10000 * 1.5 * 60 * 60 * 24
	total := 12960000000
	//查询特征 = 1.5 * 60 * 60 * 24 = 129600
	query := 129600


	now := time.Now()
	nanoStart := now.Nanosecond()
	for i:=0 ; i < total; i++{
		for j:=0; j<query; j++{
			//doNothing
		}
	}
	now = time.Now()
	nanoEnd := now.Nanosecond()
	ts := nanoEnd - nanoStart
	fmt.Println("ts(ms)=",ts/1000000," ts(nano)=",ts)
}

func main() {
	//var curDir = GetCurrentDirectory()
	//fmt.Println("current dir=", curDir)
	//
	//listen_port := flag.Int("listen_port", 80, "param for port")
	//flag.Parse()
	//fmt.Println(*listen_port)
	//addr := ":" + String(listen_port)
	////addr := ":"+string(80)
	//fmt.Println("listen port-", addr)

	testGetTime()

	//t1 := uint64(1200)
	//t2 := float32(23.976)
	//t3 := float32(t1)/t2
	//fmt.Println("---t1:",t1)
	//fmt.Println("---t2:",t2)
	//fmt.Println("---t3:",t3)

	//TestString()

	//TestGetEnv()

	//testA :=10
	//testB := testA >> 5
	//fmt.Println("testA=",testA)
	//fmt.Println("testA >> 2=",testB)

	//TestCalcute()

	//s := 0.5
	//f_s := 1 - math.Pow((1 - math.Pow(s, 28)), 20)
	//fmt.Println("s=",s," f_s=",f_s)
}
