package main

import "fmt"

/**
输出结果
[0 1 2 100 200 5 6 7 8 9]
[0 1 2]
[0 1 2 100 200]
*/
func test() {
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := data[:3]
	s2 := append(s, 100, 200) // 添加多个值
	fmt.Println(data)
	fmt.Println(s)
	fmt.Println(s2)
}

/**
输出结果
3
6 11
[0 1 100 200] [0 1 2 3 4 0 0 0 0 0 0]
0xc420016090 0xc42007a060
*/
func test1() {
	data := [...]int{0, 1, 2, 3, 4, 10: 0}
	s := data[:2:3]
	fmt.Println(cap(s))
	s = append(s, 100, 200) //append两个值，超出了s的cap，所以会重新分配底层数组，与原数组无关，通常会以两倍容量重新分配数组
	fmt.Println(cap(s), cap(data))
	fmt.Println(s, data)
	fmt.Println(&s[0], &data[0])
}

func sliceAppendTest() {
	// test()
	test1()
}
