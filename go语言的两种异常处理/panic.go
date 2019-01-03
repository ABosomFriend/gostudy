package main

import (
    "errors"
	"fmt"
)

func innerFunc() {
	fmt.Println("Enter innerFunc")
	panic(errors.New("Occur a panic!"))
	fmt.Println("Quit innerFunc")
}

func outerFunc() {
	fmt.Println("Enter outerFunc")
	innerFunc()
	fmt.Println("Quit outerFunc")
}

//interface{}代表空接口。Go语言中的任何类型都是它的实现类型
func main() {
	fmt.Println("Enter main")
	//panic相当于我们java中的运行时异常,这里的recover()函数就是捕获当前这个异常,然后recover()必须与defer
	//连用,并且要写在抛出异常的函数的前面才会生效,这里的defer我们可以类比java中的finally

	//同时这里我们也可以猜想error是一些检查异常,panic是免检异常
	defer func() {
	    if p := recover(); p != nil {
	        fmt.Printf("Fatal error: %s\n", p)
	    }
	}()
	outerFunc()
	fmt.Println("Quit main")
	
}