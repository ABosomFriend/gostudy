package main

import (
	"fmt"
)

//注意:由于字符串类型的空值为"",为空字符串

func main() {
	//字典类型的key的类型必须是可比较的
	mm2 := map[string]int{"golang": 42, "java": 1, "python": 0, "scala": 25, "erlang": 50, "test": 0}
	fmt.Println(mm2)
	//删除mm2中key为test的键值对
	delete(mm2, "test")
	fmt.Println(mm2)
	fmt.Printf("%d, %d, %d \n", mm2["scala"], mm2["erlang"], mm2["python"])
	//虽然key为test的键值对并不存在,但是还是会输出它的零值,所以这里就会存在一个歧义,到底不存在还是存的为零值
	fmt.Printf("%d\n", mm2["test"])
	//但是我们可以通过下面这种方式来判断其是否是零值,如果返回值的第二个参数为false,代表map中不存在这个key的键值对，其对应的value就为零值,为true代表map
	//中存在这个键值对
	var value, exist = mm2["test"]
	fmt.Printf("value = %d  exist = %v\n", value, exist)
	//    最后，与切片类型相同，字典类型属于引用类型。它的零值即为nil
	var map1 map[int]string = nil
	fmt.Println(map1)
}
