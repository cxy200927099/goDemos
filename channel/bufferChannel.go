package main

import "fmt"

var c = make(chan string, 1)

func f() {

	c <- "goroutine"

	fmt.Println("在goroutine内")
}

func main() {
	go f()

	//这里要说明一下，虽然channel限制容量为1，但是还是可以写多次进去的，这里不要理解错了
	c <- "main goroutine"
	fmt.Println(<-c)
	fmt.Println(<-c)

	fmt.Println("外部调用")
}

/*
运行： go run bufferChannel.go
结果：
main goroutine
goroutine
外部调用
在goroutine内

*/
