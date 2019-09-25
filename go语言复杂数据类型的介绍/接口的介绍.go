package main

import "fmt"

type Animal interface {
	Grow()
	Move(string) string
}

/*我们无需在一个数据类型中声明它实现了哪个接口。只要满足了“方法集合为其超集”的条件，
就建立了“实现”关系。这是典型的无侵入式的接口实现方法*/
type Cat struct {
	Name  string
	Age   int32
	Where string
}

func (cat *Cat) Grow() {

}

func (cat *Cat) Move(address string) string {
	return ""
}

/*
现在我们已经认为*Cat类型实现了Animal接口

在Go语言中，这种判定(某个数据类型是否为接口的实现)可以用类型断言来实现。不过，在这里，我们是不能在一个非接口类型的
值上应用类型断言来判定它是否属于某一个接口类型的。我们必须先把前者转换成空接口类型的值*/

func main() {
	myCat := Cat{"Little C", 2, "In the house"}
	/*Go语言的类型转换规则定义了是否能够以及怎样可以把一个类型的值转换另一个类型的值。
	另一方面，所谓空接口类型即是不包含任何方法声明的接口类型，用interface{}表示，
	常简称为空接口。正因为空接口的定义，Go语言中的包含预定义的任何数据类型都可以被看做是空接口的实现
	*/

	/*类型断言表达式v.(Animal)的求值结果可以有两个。第一个结果是被转换后的那个目标类型（这里是Animal）的值，
	而第二个结果则是转换操作成功与否的标志。显然，ok代表了一个bool类型的值。它也是这里判定实现关系的重要依据。*/
	animal, ok := interface{}(&myCat).(Animal)
	//%T --->打印*com.Cat
	fmt.Printf("%v, %v\n", ok, animal)
}
