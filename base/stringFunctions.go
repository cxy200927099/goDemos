/*
	string 常用的一些方法
*/
package main

import s "strings"
import "fmt"

//We alias fmt.Println to a shorter name as we’ll use it a lot below.
var p = fmt.Println

func main() {

	//Here’s a sample of the functions available in strings.
	//Since these are functions from the package, not methods on the string object itself,
	//we need pass the string in question as the first argument to the function.
	//You can find more functions in the strings package docs(http://golang.org/pkg/strings/).
	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
	p()

	//这里函数并不属于strings，但是值得提起
	p("Len: ", len("hello"))
	p("Char:", "hello"[1])
}

/*
运行： go run stringFunctions.go
结果：
Contains:   true
Count:      2
HasPrefix:  true
HasSuffix:  true
Index:      1
Join:       a-b
Repeat:     aaaaa
Replace:    f00
Replace:    f0o
Split:      [a b c d e]
ToLower:    test
ToUpper:    TEST

Len:  5
Char: 101
*/
