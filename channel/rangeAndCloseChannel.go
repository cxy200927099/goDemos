/**
之前chanel这个例子中，我们需要读取两次c，这样不是很方便，Go考虑到了这一点，
所以也可以通过range，像操作slice或者map一样操作缓存类型的channel，请看下面的例子

*/

package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
