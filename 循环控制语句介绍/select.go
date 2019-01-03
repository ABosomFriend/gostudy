package main

import "fmt"

/*如果该select语句被执行时通道ch1和ch2中都没有任何数据，那么肯定只有default case会被执行。
但是，只要有一个通道在当时有数据就不会轮到default case执行了。显然，对于包含通道接收操作的case来讲，
其执行条件就是通道中存在数据（或者说通道未空）。如果在当时有数据的通道多于一个，
那么Go语言会通过一种伪随机的算法来决定哪一个case将被执行。*/

/*另一方面，对于包含通道发送操作的case来讲，其执行条件就是通道中至少还能缓冲一个数据（或者说通道未满）。
类似的，当有多个case中的通道未满时，它们会被随机选择*/

/* 请注意，如果一条select语句中不存在default case， 并且在被执行时其中的所有case都不满足执行条件，
那么它的执行将会被阻塞！当前流程的进行也会因此而停滞。直到其中一个case满足了执行条件，执行才会继续。
我们一直在说case执行条件的满足与否取决于其操作的通道在当时的状态。这里特别强调一点，
即：未被初始化的通道会使操作它的case永远满足不了执行条件。对于针对它的发送操作和接收操作来说都是如此。
  
最后提一句，break语句也可以被包含在select语句中的case语句中。它的作用是立即结束当前的select语句的执行，
不论其所属的case语句中是否还有未被执行的语句。*/

func main() {
	ch4 := make(chan int, 1)
	//由于通道为空,所以下面的代码会发生阻塞
	// _, ok := <- ch4
	// fmt.Println(ok)
	for i := 0; i < 4; i++ {
		select {
		//但是在case这里不存在死锁这个问题,这里注意case或者default下面的语句都是他们的子语句
		case e, ok := <-ch4:
			if !ok {
				fmt.Println("End.")
				return
			}
			fmt.Println(e)
			close(ch4)
		default:
			fmt.Println("No Data!")
			ch4 <- 1
		}
	}

	ch1 := make(chan int, 0)
	select {
	case ch1 <- 123:
		fmt.Println("存入成功")
	default:
		fmt.Println("通道已满")
	}
}