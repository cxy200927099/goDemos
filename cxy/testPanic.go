package main

import (
	"net/http"
	"fmt"
	"github.com/cihub/seelog"
	"time"
	"sync"
)

//访问接口：
//curl -H 'multipart/form-data' -X POST -F "val=1" "http://127.0.0.1:80/test"

func testHandler(w http.ResponseWriter,r *http.Request){
	fmt.Println("--------debug:test msg--------")

	if r.Method!="POST" && r.Header.Get("Content-Type")!="multipart/form-data"{
		fmt.Println("post protocol error")
		fmt.Fprintf(w, "{\"code\":105, \"msg\":\"%s\"}", "post msg format error")
		return
	}

	//读取数据,错误则直接返回
	val := r.FormValue("val")		//string
	if val==""{
		seelog.Error("video name is empty:")
		fmt.Fprintf(w, "{\"code\":106, \"msg\":\"%s\"}", "video name is empty")
		return
	}

	globeAddLock.Lock()
	defer globeAddLock.Unlock()
	fmt.Println("--------debug:test msg--------val=",val)

	if val == "3"{
		panic("testPanic")
	}

	if val == "5"{
		fmt.Fprintf(w, "{\"code\":100,\"msg\":\"already exist!\"}")
		return
	}
	time.Sleep(time.Second * 2)

	fmt.Fprintf(w, "{\"code\":0,\"msg\":\"success\"}")
}

var globeAddLock sync.Mutex

func main(){
	http.HandleFunc("/test",testHandler)

	httpAddress := "127.0.0.1:8080"
	fmt.Println("http server address ",httpAddress)
	err :=http.ListenAndServe(httpAddress,nil)
	if err!=nil{
		fmt.Println("error:",err)
	}
	fmt.Println("http server end")
}