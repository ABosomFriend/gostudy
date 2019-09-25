package main

import "fmt"

func createSlice() {
	//使用这种方式创建切片
	var slice = make([]int, 10)
	fmt.Println(slice)
}
