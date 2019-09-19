package main

import "fmt"

func test1() {

	/**
		defer recover()      //无效
	    defer fmt.Println(recover())    //无效
	    defer func() {
	        func() {
	            println("defer inner")
	            recover()                //无效
	}()
			}()
	*/

	defer func() {
		fmt.Println(recover()) //其只能捕获最后抛出的panic,注意recover()函数只有在延迟调用内部直接调用才有效
	}()
	defer func() {
		panic("defer panic")
	}()
	panic("test panic")
}
func main() {
	test1() //output: defer panic
	test2()
}

//我们也阔以通过这种方式来接收panic
func recoverManager() {
	if err := recover(); err != nil {
		println(err.(string))
	}
}

func test2() {
	defer recoverManager()
	panic("panic test")
}
