package main

import "fmt"
import "sort"

func main() {

	//Sort methods are specific to the builtin type; here’s an example for strings.
	//注意Sort是改变了数组的顺序，并不会返回新的数组
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)
	//An example of sorting ints.
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)
	//We can also use sort to check if a slice is already in sorted order.
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}

/*
运行： go run sort.go
结果：
Strings: [a b c]
Ints:    [2 4 7]
Sorted:  true
*/
