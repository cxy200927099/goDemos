/*
	panic表示程序出现了意想不到的错误
	注意：在go中，经常用的是使用错误指示返回值，而不像其他语言使用exception来处理错误
*/
package main

import "os"

func main() {

	//We’ll use panic throughout this site to check for unexpected errors.
	//This is the only program on the site designed to panic.
	panic("a problem")

	//A common use of panic is to abort if a function returns an error value
	//that we don’t know how to (or want to) handle. Here’s an example of panicking
	//if we get an unexpected error when creating a new file.
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}

/*
程序运行会在11行中断，打印出相应的信息
运行： go run panic.go
结果：
panic: a problem

goroutine 1 [running]:
main.main()
        /Users/chenxingyi/work/go/code/base/panic.go:11 +0x64
exit status 2
*/
