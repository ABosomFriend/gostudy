package main

import "fmt"

func main() {
	/*
		在go语言中可以对变量进行重定义,if语句代码块中的number在其内部有效,main函数的number在外部有效(如果if代码块没有number变量
		的话,main函数中定义的number在if代码块中也是有效的)
	*/
    var number int = 5
	if number := 4; 10 > number {
		
		number += 3
		fmt.Println(number)
	} else if 10 < number {
		number -= 2
		fmt.Print(number)
	}
	fmt.Println(number)
}