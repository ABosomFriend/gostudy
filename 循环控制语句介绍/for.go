package main

import (
    "fmt"
)

func main() {
	//for语句在遍历字符串的时候,index是当前字符第一个字节所在的下标
	for i, v := range "我们123" {
		fmt.Printf("i = %d, v = %c\n", i, v)
	}

	//for在遍历通道的时候,如果通道没有数据了,for会阻塞,当通道关闭了,for语句才会结束
	
	//for在遍历map的时候,并不能保证顺序
	map1 := map[int]string{1: "Golang", 2: "Java", 3: "Python", 4: "C"}
	for i, v := range map1 {
		fmt.Printf("%d: %s\n", i, v)
	}

}