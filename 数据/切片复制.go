package main

import "fmt"

func sliceCopy() {
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := data[8:]
	s2 := data[:5]
	//后面的往前面copy
	copy(s, s2) // dst:s, src:s2
	fmt.Println(s2)
	fmt.Println(data)
}

/**
输出结果
[0 1 2 3 4]
[0 1 2 3 4 5 6 7 0 1]
*/
