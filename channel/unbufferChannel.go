package main

import (
	"fmt"
	"time"
)

func writeRoutine(test_chan chan int, value int) {

	fmt.Println("write begin")
	test_chan <- value
	fmt.Println("write end")
}

func readRoutine(test_chan chan int) {

	fmt.Println("read begin")
	<-test_chan
	fmt.Println("read end")
	return
}

func main() {

	c := make(chan int)

	x := 100

	//这里需要注意，由于channel默认都是无缓冲的，也就算会阻塞，如果阻塞了主线程，会报deadlock
	//所以无论是先read(不write)或者先write(不read) channel都需要放到goroutine去执行，否则会出现deadlock
	// readRoutine(c)
	// go writeRoutine(c, x)

	// writeRoutine(c, x)
	// go readRoutine(c)

	//下面这两种情况，不会出现deadlock，因为阻塞发生在goroutine中
	//所以无论是先在 goroutine 中read(然后不write)或者先write(然后不read)都不会影响主线程
	// go readRoutine(c)
	// writeRoutine(c, x)

	go writeRoutine(c, x)
	readRoutine(c)

	time.Sleep(time.Second * 2)
	fmt.Println(x)
}
