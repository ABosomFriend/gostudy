package main

import (
	"fmt"
)
type Person struct {
    Name    string
	Gender  string
	Age     uint8
	Address    string
}

//这里我们可以定义一个结构体的方法
func (person * Person) Move(Address string) string {
	oldAddress := person.Address
	person.Address = Address;
	return oldAddress
}

func main() {
	//Person{Name: "Robert", Gender: "Male", Age: 33} 在赋值的时候我们也可以指定名称来赋值
	p := Person{"Robert", "Male", 33, "Beijing"}
	//这里我们创建了一个匿名的结构体
	//结构体类型可以拥有若干方法（注意，匿名结构体是不可能拥有方法的）
	p1 := struct {
		Test string
		//如果我们在初始化结构体的时候不赋值,那么结构体里面的字段就是零值
	}{}

	fmt.Printf("%s\n", p1.Test)

	oldAddress := p.Move("San Francisco")
	fmt.Printf("%s moved from %s to %s.\n", p.Name, oldAddress, p.Address)

	/*最后，结构体类型属于值类型。它的零值并不是nil，而是其中字段的值均为相应类型的零值的值。
	举个例子，结构体类型Person的零值若用字面量来表示的话则为Person{}*/

	var p2 Person = Person{}
	fmt.Println(p2.Name)

}