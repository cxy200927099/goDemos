package main

import "time"
import "fmt"

func main() {

	//First we’ll look at basic rate limiting. 假设我们要限制输入请求的处理.
	//We’ll serve these requests off a channel of the same name.
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	//This limiter channel will receive a value every 200 milliseconds.
	//This is the regulator in our rate limiting scheme.
	limiter := time.Tick(time.Millisecond * 200)

	//By blocking on a receive from the limiter channel before serving each request,
	//we limit ourselves to 1 request every 200 milliseconds.
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	fmt.Println("=====================")
	time.Sleep(time.Second * 2)

	//We may want to allow short bursts of requests in our rate limiting scheme
	//while preserving the overall rate limit. We can accomplish this by buffering
	// our limiter channel. This burstyLimiter channel will allow bursts of up to 3 events.
	burstyLimiter := make(chan time.Time, 3)

	//Fill up the channel to represent allowed bursting.
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	//Every 200 milliseconds we’ll try to add a new value to burstyLimiter, up to its limit of 3.
	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	//Now simulate 5 more incoming requests. The first 3 of these will
	//benefit from the burst capability of burstyLimiter.
	const REQUEST_COUNT int = 5
	burstyRequests := make(chan int, REQUEST_COUNT)
	for i := 1; i <= REQUEST_COUNT; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}

/*
Our running program shows the 5 jobs being executed by various workers.
尽管这个工作量是5秒，但是我们的程序却只花费了2秒，因为有三个worker在并行的执行
运行： go run workerPools.go
结果：
//可以看到第一批处理是按照我们所期望的每隔200ms去处理的
request 1 2017-11-10 17:11:15.063645246 +0800 CST
request 2 2017-11-10 17:11:15.265627252 +0800 CST
request 3 2017-11-10 17:11:15.46357249 +0800 CST
request 4 2017-11-10 17:11:15.663766117 +0800 CST
request 5 2017-11-10 17:11:15.864654917 +0800 CST

//第二批处理的前三个几乎是同时处理，然后后面是每隔200ms处理剩下的
request 1 2017-11-10 17:11:15.864728676 +0800 CST
request 2 2017-11-10 17:11:15.864736582 +0800 CST
request 3 2017-11-10 17:11:15.864742954 +0800 CST
request 4 2017-11-10 17:11:16.065613423 +0800 CST
request 5 2017-11-10 17:11:16.265100062 +0800 CST
*/
