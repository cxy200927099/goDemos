/**
* 匿名函数，闭合函数
 */
package main

import "fmt"

//定义一个函数，这个函数的返回值也是一个函数 func() int
func intSeq() func() int {

	i := 0
	//这里返回的函数就称作匿名函数
	return func() int {
		i++
		return i
	}
}

func main() {

	//定义一个nextInt函数指针，给其赋值
	nextInt := intSeq()
	//调用的时候与普通函数的调用是一样的
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInt := intSeq()
	fmt.Println(newInt())
}
