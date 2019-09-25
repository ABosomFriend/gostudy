package main

import "fmt"

//User 用户信息
//User结构体的属性小写，只能是在包中可见
type User struct {
	age  int
	name string
}

func main() {
	user := User{23, "yizijun"}
	fmt.Println(user)
}
