package main

import "fmt"

func arrayTest() {
	var array1 = [...]int{1, 2, 3}
	fmt.Printf("%d\n", array1[2])
	//[1,2,3]
	fmt.Println(array1)
	//打印这种boolean类型用%t
	fmt.Printf("%t\n", true)
	fmt.Printf("array1.size = %d\n", len(array1))
	//默认都会为每个index赋值为零值
	var array2 [5]int
	fmt.Printf("len(array2) = %d array2[0]= %d\n", len(array2), array2[0])
}

/**
输出结果：
3
[1 2 3]
true
array1.size = 3
len(array2) = 5   array2[0]= 0
*/
