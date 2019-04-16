/**
 * goroutine是Go并行设计的核心。goroutine说到底其实就是协程，但是它比线程更小，十几个goroutine可能体现在底层就是五六个线程，Go语言内部帮你实现了这些goroutine之间的内存共享。
 执行goroutine只需极少的栈内存(大概是4~5KB)，当然会根据相应的数据伸缩。也正因为如此，可同时运行成千上万个并发任务。goroutine比thread更易用、更高效、更轻便。

goroutine是通过Go的runtime管理的一个线程管理器。goroutine通过go关键字实现了，其实就是一个普通的函数。

    go hello(a, b, c)
通过关键字go就启动了一个goroutine。我们来看一个例子
*/

/**
runtime.Gosched()表示让CPU把时间片让给别人,下次某个时候继续恢复执行该goroutine。

默认情况下，在Go 1.5将标识并发系统线程个数的runtime.GOMAXPROCS的初始值由1改为了运行环境的CPU核数。
*/
package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		//
		runtime.Gosched()
		fmt.Println(s)
	}
}

func main() {
	go say("world") //开一个新的Goroutines执行
	say("hello")    //当前Goroutines执行
	fmt.Println("CPU number: ", runtime.GOMAXPROCS(1000))
	fmt.Println(runtime.GOMAXPROCS(1000))
	fmt.Println(runtime.GOMAXPROCS(-1))
}
