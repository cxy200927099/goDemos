/*
结构体
*/
package main

import "fmt"

//结构体的定义
type person struct {
	name string
	age  int
}

func main() {

	//创建新的结构体
	fmt.Println(person{"Bob", 20})
	//这种情况创建结构体，指定了名称
	fmt.Println(person{name: "Alice", age: 30})
	//只给一些属性赋值
	fmt.Println(person{name: "Fred"})
	//结构体取地址
	fmt.Println(&person{name: "Ann", age: 40})
	//结构体赋值给一个变量
	s := person{name: "Sean", age: 50}
	fmt.Println(s)
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp)
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

}

/*
运行: go run Struct.go
输出结果:
{Bob 20}
{Alice 30}
{Fred 0}
&{Ann 40}
{Sean 50}
Sean
&{Sean 50}
50
51
*/
