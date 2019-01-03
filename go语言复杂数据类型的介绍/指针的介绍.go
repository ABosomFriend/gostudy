package main

import "fmt"

type MyInt struct {
    n int
}

type Pet interface {
    Name() string
	Age() uint8
}

type Dog struct {
	FName string
	FAge uint64
}

func (dog Dog) Name() string {
  return ""
}

func (dog Dog) Age() uint8 {
  return 1
}

//这里的方法接收者是一个指针类型,说明这个方法是属于*MyInt的.
func (myInt *MyInt) Increase() {
	//下面代码是(*)myInt.n++的语法糖
	myInt.n++
}

//这个方法的接收者是MyInt类型
func (myInt MyInt) Decrease() {
	//但是下面的自减操作并不会成功.因为这里是值传递,并不是址传递,我们可以类比c语言中的指针
	myInt.n--
}

/*
相关规则:
1.一个指针类型拥有以它以及以它的基底类型为接收者类型的所有方法，而它的基底类型却只拥有以它本身为接收者类型的方法。
*/
func main() {
	/*另外，还有一点需要大家注意，我们在基底类型的值上仍然可以调用它的指针方法。
	例如，若我们有一个Person类型的变量bp，则调用表达式bp.Grow()是合法的。这是因为，
	如果Go语言发现我们调用的Grow方法是bp的指针方法，那么它会把该调用表达式视为(&bp).Grow()。
	实际上，这时的bp.Grow()是(&bp).Grow()的速记法*/
	mi := MyInt{}
	//比如像下面的表达式是(&mi).Increase()的语法糖
	mi.Increase()
	mi.Increase()
	mi.Decrease()
	mi.Decrease()
	mi.Increase()
	fmt.Println(mi.n)
	fmt.Printf("%v\n", mi.n == 1)

	myDog := Dog{"Little D", 3}
	//了解下面这种写法
	_, ok1 := interface{}(&myDog).(Pet)
	_, ok2 := interface{}(myDog).(Pet)
	//如果上面两个实现方法的输入值都是Dog类型,那么Dog 和 *Dog都是Pet接口的实现类型
	fmt.Printf("%v, %v\n", ok1, ok2)
}