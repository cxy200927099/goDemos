package main

import "time"
import "fmt"

func main() {
	//Tickers use a similar mechanism to timers: a channel that is sent values.
	//Here we’ll use the range builtin on the channel to iterate over the values
	// as they arrive every 500ms.
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()
	//Tickers can be stopped like timers. Once a ticker is stopped
	// it won’t receive any more values on its channel. We’ll stop ours after 1600ms.
	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}

/*
运行： go run ticker.go
结果：
Tick at 2017-11-10 15:10:22.205519318 +0800 CST
Tick at 2017-11-10 15:10:22.701988875 +0800 CST
Tick at 2017-11-10 15:10:23.202038129 +0800 CST
Ticker stopped

*/
