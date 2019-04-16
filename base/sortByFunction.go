/*
	有时候我们想按照其他的方式排序，而不是按照字母，数字大小之类的，
	例如我们想按照字符串长度去排序一个字符串，这里就给出了一个自定义排序的例子
*/
package main

import "sort"
import "fmt"

//In order to sort by a custom function in Go, we need a corresponding type.
//这里我们为内置的p[]string定义一个别名
type ByLength []string

//We implement sort.Interface - Len, Less, and Swap -
//on our type so we can use the sort package’s generic Sort function.
//Len and Swap will usually be similar across types
//and Less will hold the actual custom sorting logic.
//In our case we want to sort in order of increasing string length,
//so we use len(s[i]) and len(s[j]) here.
func (s ByLength) Len() int {
	return len(s)
}
func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

//With all of this in place, we can now implement our custom sort by
//casting the original fruits slice to ByLength,
//and then use sort.Sort on that typed slice.
func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}

/*
运行： go run sort.go
结果：
[kiwi peach banana]
*/
