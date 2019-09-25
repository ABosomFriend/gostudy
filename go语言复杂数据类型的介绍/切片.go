package main

import "fmt"

func main() {
	var number1 = []int{1, 1, 1, 1}
	fmt.Printf("%d\n", number1[0])
	var numbers3 = [5]int{1, 2, 3, 4, 5}
	//slice3的结果为[]int{3,4,5}
	//只包括下界,不包括上界
	slice3 := numbers3[2:len(numbers3)]
	fmt.Println(slice3)
	length := (3)
	capacity := (3)
	//一个切片值的容量即为它的第一个元素值在其底层数组中的索引值与该数组长度的差值的绝对值
	fmt.Printf("%t, %t\n", (length == len(slice3)), (capacity == cap(slice3)))
	/*最后，要注意，切片类型属于引用类型。它的零值即为nil，即空值。
	  如果我们只声明一个切片类型的变量而不为它赋值，那么该变量的值将会是nil*/
	var slice4 []string
	fmt.Println(slice4)
	var numbers4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//[]int{5,6}
	slice5 := numbers4[4:6:8]
	//cap(slice5) = 4 因为上面限制了容量的大小为8,所以cap(slice5) = 8 - 4
	fmt.Printf("cap(slice5) = %d\n", cap(slice5))
	// slice6 := numbers4[4:5:5]
	//下面的语句报错,因为slice6切片的上界不能大于等于第三个参数的大小,第三个参数表示切片的容量
	// fmt.Printf("slice6[6] = %d\n", slice6[5])
	length1 := (2)
	capacity1 := (4)
	fmt.Printf("%t, %t\n", length1 == len(slice5), capacity1 == cap(slice5))
	//cap(slice5) == 4 slice5现在变成 []int{5,6,7,8}
	//我们可以通过如下操作将其长度延展得与其容量相同：
	slice5 = slice5[:cap(slice5)]
	fmt.Println(slice5)
	//slice5 []int{5,6,7,8,11,12,13}, append函数是返回一个新的切片
	slice5 = append(slice5, 11, 12, 13)
	length1 = (7)
	fmt.Printf("%t\n", length1 == len(slice5))
	fmt.Printf("slice5 cap = %d\n", cap(slice5))
	slice7 := []int{0, 0, 0}
	/*1. 这种复制遵循最小复制原则，即：被复制的元素的个数总是等于长度较短的那个参数值的长度。
	  2. 与append函数不同，copy函数会直接对其第一个参数值进行修改。*/
	copy(slice5, slice7)
	e2 := (0)
	e3 := (8)
	e4 := (11)
	fmt.Printf("%t, %t, %t\n", e2 == slice5[2], e3 == slice5[3], e4 == slice5[4])
	var slice8 = make([]int, 0)
	//false
	fmt.Printf("slice8 == nil? %t\n", slice8 == nil)
	//等同于完整的数组
	var slice9 = numbers4[:]
	fmt.Println(slice9)
}

/**
输出结果:
1
[3 4 5]
true, true
[]
cap(slice5) = 4
true, true
[5 6 7 8]
true
slice5 cap = 8
true, true, true
slice8 == nil? false
[1 2 3 4 5 6 7 8 9 10]

*/
