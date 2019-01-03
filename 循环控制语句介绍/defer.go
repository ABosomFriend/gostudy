package main

import (
    "fmt"
)
/*func readFile(path string) ([]byte, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    return ioutil.ReadAll(file)
}

 defer所指向的语句在方法结束之前会执行,不论是正常返回,还是抛出异常*/
func main() {
	for i := 0; i < 10; i++ {
		defer func(num int) {
		    fmt.Printf("%d ", num)
		}(func () int {
			num := fibonacci(i)
			fmt.Printf("%d ", num)
		    return num
		}())
	}
	// deferIt()
	// deferIt3()
	// deferIt4()
}

//当一个函数中存在多个defer语句时，它们携带的表达式语句的执行顺序一定是它们的出现顺序的倒序
func deferIt() {
    defer func() {
        fmt.Print(1)
    }()
    defer func() {
        fmt.Print(2)
    }()
    defer func() {
        fmt.Print(3)
    }()
    fmt.Print(4)
}

//执行结果  1 2 3 4 40 30 20 10
/*defer携带的表达式语句代表的是对某个函数或方法的调用。这个调用可能会有参数传入，
比如：fmt.Print(i + 1)。如果代表传入参数的是一个表达式，那么在defer语句被执行的
时候该表达式就会被求值了。注意，这与被携带的表达式语句的执行时机是不同的。请揣测下面这段代码的执行*/
func deferIt3() {
    f := func(i int) int {
        fmt.Printf("%d ",i)
        return i * 10
    }
    for i := 1; i < 5; i++ {
        defer fmt.Printf("%d ", f(i))
    }
}

func fibonacci(num int) int {
	if num == 0 {
		return 0
	}
	if num < 2 {
		return 1
	}
	return fibonacci(num-1) + fibonacci(num-2)
}

/*deferIt4函数在被执行之后标出输出上会出现5555，而不是4321。
原因是defer语句携带的表达式语句中的那个匿名函数包含了对外部（确切地说，是该defer语句之外）的变量的使用。
注意，等到这个匿名函数要被执行（且会被执行4次）的时候，包含该defer语句的那条for语句已经执行完毕了。
此时的变量i的值已经变为了5。因此该匿名函数中的打印函数只会打印出5。正确的用法是：把要使用的外部变量
作为参数传入到匿名函数中。修正后的deferIt4函数如下：

func deferIt4() {
    for i := 1; i < 5; i++ {
        defer func(n int) {
            fmt.Print(n)
        }(i)
    }
}

*/
func deferIt4() {
    for i := 1; i < 5; i++ {
        defer func() {
            fmt.Print(i)
        }()
    }
}