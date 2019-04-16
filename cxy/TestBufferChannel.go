package main

import (
	"fmt"
	"time"
	"math"
)

func handler(c chan int){

	fmt.Println("ready to read data")
	for {
		readData := <- c
		fmt.Println("read data:",readData)
	}

}


func main(){

	B := 1
	KB := 1024 * B
	MB := 1024 * KB
	GB := 1024 * MB
	//10GB
	test10G := 10 * GB
	fmt.Println("10GB=",test10G)
	ta := float64(940/1024.0)
	inta := 10
	size := uint64( math.Floor( float64(940/1024.0) + 0.5 ) )
	fmt.Println("float64(940/1024.0)=",ta," inta=",inta, " size=",size)

	channel := make(chan int, 10)
	//o := make(chan bool)
	go func() {
		handler(channel)
	}()

	fmt.Println("prepare for write")
	time.AfterFunc(1*time.Second, func() {

		fmt.Println("ready to write")
		for i := 0; i < 5; i++ {
			fmt.Println("write data:",i)
			channel <- i
			time.Sleep(1*time.Second)
		}
	})
	time.Sleep(10*time.Second)

}


