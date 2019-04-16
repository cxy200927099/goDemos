/*
给结构体赋值方法
*/
package main

import "fmt"

//结构体的定义
type rect struct {
	width  int
	height int
}

//give rect a method area()
func (r *rect) area() int {
	return r.width * r.height
}

//another way to give rect a method perim()
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {

	//结构体赋值给一个变量
	r := rect{width: 10, height: 3}
	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())

}

/*
运行: go run structMethod.go
输出结果:
area:  30
perim:  26
area:  30
perim:  26
*/
