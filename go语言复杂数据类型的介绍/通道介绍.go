package main

import (
	"fmt"
	"time"
)

/*我们还可以以数据在通道中的传输方向为依据来划分通道。默认情况下，通道都是双向的，即双向通道。
如果数据只能在通道中单向传输，那么该通道就被称作单向通道。我们在初始化一个通道值的时候不能指定它为单向
但是我们可以通过下面的方式来设置单向通道*/
type Sender chan<- int

type Receiver <-chan int

/*
关于通道的其它几点性质
1.通道不能多次close操作,否则要抛出运行时异常
2.通道是基于FIFO的操作
3.在通道有效的前提下,如果通道已满,再向通道中发送数据,则会阻塞
4.当通道被关闭了后,再往通道中发送数据,则会抛出运行时异常
5.如果通道已为空,这个时候接受者取数据也会被阻塞
*/

/*通道（Channel）是Go语言中一种非常独特的数据结构。它可用于在不同Goroutine之间传递类型化的数据，并且是并发安全的。
相比之下，我们之前介绍的那些数据类型都不是并发安全的*/

//顺便说一句，实际上make函数也可以被用来初始化切片类型或字典类型的值。
func main() {
	//var ch2 chan string = make(chan string, 5)
	ch2 := make(chan string, 5)
	// 下面就是传说中的通过启用一个Goroutine来并发的执行代码块的方法。
	// 关键字 go 后跟的就是需要被并发执行的代码块，它由一个匿名函数代表。
	// 对于 go 关键字以及函数编写方法，我们后面再做专门介绍。
	// 在这里，我们只要知道在花括号中的就是将要被并发执行的代码就可以了。
	go func() {
		ch2 <- ("向通道发送字符串")
	}()

	var value string = "数据"
	//关闭通道
	close(ch2)
	/*这样做的目的同样是为了消除与零值有关的歧义。这里的变量ok的值同样是bool类型的。
	它代表了通道值的状态，true代表通道值有效，而false则代表通道值已无效（或称已关闭）,通道关闭了就会返回一个零值*/
	//var out string = <- ch2
	var out, ok = <-ch2
	fmt.Printf("ok?%v\n", ok)
	value = value + out
	fmt.Println(value)
	//不能多次关闭通道,多次关闭通道要抛出运行时异常
	// close(ch2)


	//当我们把缓冲区设置为0的时候,这个时候通道就没有缓存.发送数据时会直接阻塞,知道接收者取出数据为止
	var myChannel = make(chan int,0)
	var number = 6
	go func() {
		var sender Sender = myChannel
		sender <- number
		//不能把单向通道赋值给双向通道
		//var chanle chan int = sender
		fmt.Println("Sent!")
	}()
	go func() {
		var receiver Receiver = myChannel
		fmt.Println("Received!", <-receiver)
	}()
	// 让main函数执行结束的时间延迟1秒，
	// 以使上面两个代码块有机会被执行。
	time.Sleep(time.Second)
}
