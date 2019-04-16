package main

import "fmt"

func main() {
	//We’ll iterate over 2 values in the queue channel.
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"

	//这里关闭了channel，但是任然可以使用range从channel中读取数据
	close(queue)
	//This range iterates over each element as it’s received from queue.
	//Because we closed the channel above, the iteration terminates after receiving the 2 elements.
	//如果不close channel，会报错deadlock
	for elem := range queue {
		fmt.Println(elem)
	}
}

/*
这个例子也说明了，可以关闭一个非空的channel，但是保留了channel的数据

运行：go run rangeoverChannel.go
结果：
one
two

*/
