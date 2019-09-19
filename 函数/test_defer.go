package main

func test(x int) {

	//这里还是会输出 c b a
	defer println("a")
	defer println("b")
	defer func() {
		println(100 / x) //div 0异常未被捕获，逐渐向外传递。最终终止进程
	}()
	defer println("c")
}
func main() {
	test(0)
}
