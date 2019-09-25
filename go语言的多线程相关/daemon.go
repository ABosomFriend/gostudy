package main

import (
	"fmt"
	"time"
)

//默认是daemon线程
func main() {
	go func() {
		for i := 0; i < 10000; i++ {
			fmt.Println(i)
		}
	}()
	time.Sleep(time.Nanosecond * 100000)
	fmt.Println("end of main")
	fmt.Println(time.Now().UnixNano())
}

/**
代码执行结果：
….
end of main
.....
112
113
114
115
116
117
118
119
1531560933145781437
*/
