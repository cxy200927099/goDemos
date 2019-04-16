/*
通过引入包sync/atomic的atomic counters来解决多个goroutine访问同一资源
由于每个携程每次运行都是系统去协调的，不能保证每一次运行的结果都一样
*/
package main

import "fmt"
import "time"
import "sync/atomic"

func main() {
	//We’ll use an unsigned integer to represent our (always-positive) counter.
	var ops uint64 = 0
	//To simulate concurrent updates, we’ll start 50 goroutines
	//that each increment the counter about once a millisecond.
	for i := 0; i < 50; i++ {
		go func() {
			for {
				//To atomically increment the counter we use AddUint64,
				//giving it the memory address of our ops counter with the & syntax.
				atomic.AddUint64(&ops, 1)

				//Wait a bit between increments.
				time.Sleep(time.Millisecond)
			}
		}()
	}

	//等待一分钟，让这个50个携程去执行加1的操作
	time.Sleep(time.Second)
	// time.Sleep(time.Millisecond * 100)

	//In order to safely use the counter while it’s still being updated
	//by other goroutines, we extract a copy of the current value
	//into opsFinal via LoadUint64. As above we need to give this function
	// the memory address &ops from which to fetch the value.
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}

/*
这里执行的结果每次都是不一样的，以下是我执行了三次的结果
运行： go run atomicCounter.go
结果：
ops: 41000

ops: 40951

ops: 40950
*/
