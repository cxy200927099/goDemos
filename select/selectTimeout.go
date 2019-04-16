package main

import (
	_ "fmt"
	"time"
)

func main() {
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			println("before select")
			select {
			case v := <-c:
				println(v)
			case <-time.After(5 * time.Second):
				println("timeout")
				o <- true
				break
			}
		}
	}()
	<-o
}

/*
运行： go run selectTimeout.go
结果：
//因为channel c没有数据写入，所以读取会一直阻塞，
//过了5秒后，time.Afte写入数据，打印了 timeout,往 o 中写入true
//主线程收到数据，退出

*/
