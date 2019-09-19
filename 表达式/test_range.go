package main

import "fmt"

/**
结果输出
[0 999 999]
[100 101 102]
*/
func testCopyObject(a [3]int) {
	for i, v := range a {
		if i == 0 {
			a[1], a[2] = 999, 999
			fmt.Println(a)
		}
		//这里的v是复制的副本
		a[i] = v + 100
	}
	fmt.Println(a)
}

/**
输出
0 1
1 2
2 100
3 4
4 5
*/
func testCopySlice(s []int) {
	for i, v := range s {
		if i == 0 {
			s = s[:3] //对slice的修改不会影响到range
			s[2] = 100
		}
		println(i, v)
	}
}

func main() {
	a := [3]int{0, 1, 2}
	testCopyObject(a)
	s := []int{1, 2, 3, 4, 5}
	testCopySlice(s)

	var slice []int
	//true
	println(slice == nil)
}
