package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//Reading files requires checking most calls for errors.
//This helper will streamline our error checks below.
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {

	//Perhaps the most basic file reading task is slurping a file’s
	//entire contents into memory.
	dat, err := ioutil.ReadFile("/tmp/dat")
	check(err)
	fmt.Print(string(dat))

	//You’ll often want more control over how and what parts of
	//a file are read. For these tasks, start by Opening a file
	//to obtain an os.File value.
	f, err := os.Open("/tmp/dat")
	check(err)

	//Read some bytes from the beginning of the file.
	//Allow up to 5 to be read but also note how many actually were read.
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	//You can also Seek to a known location in the file and Read from there.
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	//The io package provides some functions that may be helpful for
	//file reading. For example, reads like the ones above can be
	//more robustly implemented with ReadAtLeast.
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	//There is no built-in rewind, but Seek(0, 0) accomplishes this.
	_, err = f.Seek(0, 0)
	check(err)

	//The bufio package implements a buffered reader that may be useful
	//both for its efficiency with many small reads and because of
	//the additional reading methods it provides.
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	//Close the file when you’re done (usually this would be scheduled immediately after Opening with defer).
	f.Close()
}

/*
结果：
➜  go go run code/advance/fileRead.go
panic: open /tmp/dat: no such file or directory

goroutine 1 [running]:
main.check(0x111c1c0, 0xc4200141e0)
        /Users/chenxingyi/work/go/code/advance/fileRead.go:15 +0x4a
main.main()
        /Users/chenxingyi/work/go/code/advance/fileRead.go:23 +0xa9
exit status 2

因此我们需要先创建文件:"/tmp/dat",向文件中写入hello go,再次运行代码，结果如下：
➜  go go run code/advance/fileRead.go
hello go5 bytes: hello
2 bytes @ 6: go
2 bytes @ 6: go
5 bytes: hello

*/
