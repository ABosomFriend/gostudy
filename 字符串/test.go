package main

import "fmt"

func test() {
	var slice1 = make([]int, 3)
	slice1[0] = 1
	fmt.Println(slice1)

	var slice2 = new([]int)
	slice2 = &slice1
	fmt.Println(*slice2)

	var str = `abc中国`

	for i := 0; i < len(str); i++ {
		//这里是每个字节每个字节打印
		fmt.Println(str[i])
	}

	for _, val := range str {
		//这里是每个字符每个字符打印
		fmt.Printf("%c\n", val)
	}

	/**
	子串依然指向原字节数组，仅修改指针和长度属性
	*/
	var str1 = str[1:]
	fmt.Println(str1)

	var char = 'a'
	fmt.Printf("char type is %T\n", char)

	//字符串默认是不能够修改的，我们可以通过这种方式修改字符串
	var strArray = []rune(str)
	strArray[0] = 'c'
	fmt.Println(string(strArray))

}

func main() {
	test()
}
