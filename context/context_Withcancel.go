package main

import (
	"context"
	"fmt"
	"time"
)

func gen(ctx context.Context) <-chan int{
	dst := make(chan int)
	n := 1
	go func(){
		for{
			select {
			case <-ctx.Done():
				fmt.Println("e exited")
			return // returning not leak the goroutine
			case dst <- n:
				n++
			}
		}
	}()

	return dst
}

func test()  {
	// gen generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the context once
	// they are done consuming generated integers not to leak
	// the internal goroutine started by gen.
	ctx, cancel := context.WithCancel(context.Background())
	//调用cancel之后，会有context.Done
	defer cancel() // cancel when we are finished consuming integers
	intChan := gen(ctx)
	for n := range intChan {
		fmt.Println(n)
		if n == 10 {
			break
		}
	}
}

func main()  {
	test()
	time.Sleep(time.Minute)
}

//程序运行结果
/*
1
2
3
4
5
6
7
8
9
10
e exited

//大约一分钟之后，进程退出，打印
Process finished with exit code 0

 */