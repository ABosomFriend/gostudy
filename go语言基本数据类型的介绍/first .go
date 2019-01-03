package main

import (
	"fmt"
)

func main() {
	var num1 int
	num1 = 2
	var num2, num3 int = 3, 4
	var num4 int8 = -128
	const ACK string = "我们"
	//8  汉字的utf-8编码占3个字节
	fmt.Printf("%d\n", len(ACK))
	fmt.Printf("%d %d %d\n", num1, num2, num3)
	fmt.Printf("%d\n", num4)
	fmt.Printf(ACK + "\n")
	// 短变量声明语句，由变量名size、特殊标记:=，以及值（需要你来填写）组成。
	size := (8)
	fmt.Println(size)
	//如果是无符号整数，它的所有位数都是用来表示数字的大小
	var num5 uint8 = 255
	fmt.Println(num5)
	var num6 uint32 = 0x1111
	//%X表示将数字用16进制表示出来
	fmt.Printf("%x     %d\n", num6, num6)
	// 可以在变量声明并赋值的语句中，省略变量的类型部分。

	// 不过别担心，Go语言可以推导出该变量的类型。
	var num7 = 5.89E-4
	// 这里用到了字符串格式化函数。其中，%E用于以带指数部分的表示法显示浮点数类型值，%f用于以通常的方法显示浮点数类型值。
	fmt.Printf("浮点数 %E 表示的是 %f。\n", num7, (num7))

	//这里num8是一个复数类型
	var num8 complex64 = 3.7E+1 + 5.98E-2i

	// 这里用到了字符串格式化函数。其中，%E用于以带指数部分的表示法显示浮点数类型值，%f用于以通常的方法显示浮点数类型值。
	fmt.Printf("复数 %E 表示的是 %f。\n", num8, num8)

	// 声明一个rune类型变量并赋值
	var char1 rune = '赞'

	// 这里用到了字符串格式化函数。其中，%c用于显示rune类型值代表的字符。
	fmt.Printf("字符 '%c' 的Unicode代码点是 %s。\n", char1, ("U+8D5E"))

	//这里的byte类型表示我们的uint8
	var byte1 byte = 127
	fmt.Printf("%d\n", byte1)

	// 声明一个string类型变量并赋值
	var str1 string = "\\\""

	// 这里用到了字符串格式化函数。其中，%q用于显示字符串值的表象值并用双引号包裹。
	fmt.Printf("%q\n", str1)
	//2
	fmt.Printf("str1变量的长度为  %d\n", len(str1))
	fmt.Printf("用解释型字符串表示法表示的 %s 所代表的是 %s。\n", str1, `str1`)
	/*最后要注意，字符串值是不可变的。也就是说，我们一旦创建了一个此类型的值，就不可能再对它本身做任何修改。*/
	var str2 string = `\r\n`
	fmt.Printf("%s\n", str2)
	//4 \r\n  这里的len表示的为其所占的字节数,由于这里是用utf-8编码格式的,所以这里所占的长度为4个字节
	fmt.Printf("%d\n", len(str2))
	var isTrue bool = true
	//格式化输出bool类型我们可以用%t
	fmt.Printf("isTrue? %t", isTrue)

}
