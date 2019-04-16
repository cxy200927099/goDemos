/*
	string 格式化
*/
package main

import "fmt"
import "os"

type point struct {
	x, y int
}

func main() {

	//Go offers several printing “verbs” designed to format general Go values.
	//For example, this prints an instance of our point struct.
	p := point{1, 2}
	fmt.Printf("%v\n", p)

	//如果是结构体，会一起打印结构体的值和属性名称
	fmt.Printf("%+v\n", p)

	//The %#v 打印go中语法表示,
	//i.e. the source code snippet that would produce that value.
	fmt.Printf("%#v\n", p)

	//%T.打印类型
	fmt.Printf("%T\n", p)

	//格式化布尔值.
	fmt.Printf("%t\n", true)

	//%d 格式化十进制数.
	fmt.Printf("%d\n", 123)

	//%b格式化为二进制.
	fmt.Printf("%b\n", 14)

	//%c格式化为字符.
	fmt.Printf("%c\n", 33)

	//%x 格式化为16进制.
	fmt.Printf("%x\n", 456)

	//There are also several formatting options for floats. For basic decimal formatting use %f.
	fmt.Printf("%f\n", 78.9)

	//%e and %E format the float in (slightly different versions of) scientific notation.
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)

	// %s基本的字符串格式化.
	fmt.Printf("%s\n", "\"string\"")

	// %q.把双引号引用的一起打印
	fmt.Printf("%q\n", "\"string\"")

	//As with integers seen earlier, %x renders the string in base-16, with two output characters per byte of input.
	fmt.Printf("%x\n", "hex this")

	// %p.打印结构体指针地址
	fmt.Printf("%p\n", &p)

	//When formatting numbers you will often want to control the width and precision of the resulting figure. To specify the width of an integer, use a number after the % in the verb. By default the result will be right-justified and padded with spaces.
	fmt.Printf("|%6d|%6d|\n", 12, 345)

	//You can also specify the width of printed floats, though usually you’ll also want to restrict the decimal precision at the same time with the width.precision syntax.
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	//To left-justify, use the - flag.
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	//You may also want to control width when formatting strings, especially to ensure that they align in table-like output. For basic right-justified width.
	fmt.Printf("|%6s|%6s|\n", "foo", "b")

	//To left-justify use the - flag as with numbers.
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	//Sprintf formats and returns a string without printing it anywhere.
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	//You can format+print to io.Writers other than os.Stdout using Fprintf.
	fmt.Fprintf(os.Stderr, "an %s\n", "error")

	str1 := "/dev/ssss"
	fmt.Println("str1[0:3]=",str1[0:4])
}

/*
运行： go run stringFormate.go
结果：
{1 2}
{x:1 y:2}
main.point{x:1, y:2}
main.point
true
123
1110
!
1c8
78.900000
1.234000e+08
1.234000E+08
"string"
"\"string\""
6865782074686973
0xc42000e260
|    12|   345|
|  1.20|  3.45|
|1.20  |3.45  |
|   foo|     b|
|foo   |b     |
a string
an error
*/
