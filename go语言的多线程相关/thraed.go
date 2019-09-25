package main

import (
	"fmt"
	"sync"
)

/*在go语句被执行时，其携带的函数（也被称为go函数）以及要传给它的若干参数（如果有的话）
会被封装成一个实体（即Goroutine），并被放入到相应的待运行队列中。Go语言的运行时系统会
适时的从队列中取出待运行的Goroutine并执行相应的函数调用操作。注意，对传递给这里的函数
的那些参数的求值会在go语句被执行时进行。这一点也是与defer语句类似的。*/

func testSync() {
	/*
		sync.WaitGroup类型有三个方法可用——Add、Done和Wait。Add会使其所属值的一个内置整数得到相应增加，
		Done会使那个整数减1，而Wait方法会使当前Goroutine（这里是运行main函数的那个Goroutine）阻塞直到
		那个整数为0。这下你应该明白上面这个示例所采用的方法了。我们在main函数中启用了三个Goroutine来封装
		三个go函数。每个匿名函数的最后都调用了wg.Done方法，并以此表达当前的go函数会立即执行结束的情况。
		当这三个go函数都调用过wg.Done函数之后，处于main函数最后的那条wg.Wait()语句的阻塞作用将会消失，
		main函数的执行将立即结束。
	*/
	var wg, wg1, wg2 sync.WaitGroup
	wg.Add(1)
	wg1.Add(1)
	wg2.Add(1)
	go func() {
		fmt.Println("1")
		wg.Done()
	}()
	go func() {
		wg.Wait()
		fmt.Println("2")
		wg1.Done()
	}()
	go func() {
		wg1.Wait()
		fmt.Println("3")
		wg2.Done()
	}()
	/*Gosched yields the processor, allowing other goroutines to run. It does not
	  suspend the current goroutine, so execution resumes automatically.*/
	//runtime.Gosched()
	wg2.Wait()
}
