package main

import (
    "fmt"
	"math/rand"
)

func main() {
	//switch后面的值不仅可以是基本类型,还可以是type类型
	ia := []interface{}{byte(6), 'a', uint(10), int32(-4)}
	switch v := ia[rand.Intn(4)]; v.(type) {
	case int32 :
		fmt.Printf("Case A.")
	case byte :
		fmt.Printf("Case B.")
	default:
		fmt.Println("Unknown!")
	}
}