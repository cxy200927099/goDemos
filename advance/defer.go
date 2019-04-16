/*
	defer是go中用于延迟执行函数的关键字，通常用于数据销毁，关闭文件等待
	他通常会在main函数的最后才会被执行
*/
package main

import "fmt"
import "os"

//Suppose we wanted to create a file, write to it, and then close when we’re done.
//Here’s how we could do that with defer.
func main() {

	//Immediately after getting a file object with createFile,
	//we defer the closing of that file with closeFile.
	//This will be executed at the end of the enclosing function (main), after writeFile has finished.
	f := createFile("/tmp/defer.txt")
	f1 := createFile("/tmp/defer1.txt")

	//注意这里，最后会先执行closeFile(f1)，才会执行closeFile(f)
	defer closeFile(f)
	defer closeFile(f1)
	writeFile(f)

	fmt.Println("end of main")
}
func createFile(p string) *os.File {
	fmt.Println("creating file:", p)
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}
func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}
func closeFile(f *os.File) {
	fmt.Println("closing file:", f.Name())
	f.Close()
}

/*

运行： go run defer.go
结果：
creating file: /tmp/defer.txt
creating file: /tmp/defer1.txt
writing
end of main
closing file: /tmp/defer1.txt
closing file: /tmp/defer.txt
*/
