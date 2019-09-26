package main

import "fmt"

type Data struct {
	x int
}

func (self Data) ValueTest() { // func ValueTest(self Data);
	fmt.Printf("Value: %p\n", &self)
}
func (self *Data) PointerTest() {
	fmt.Printf("Pointer: %p\n", self)
}
func main() {
	d := Data{}
	p := &d
	fmt.Printf("Data: %p\n", p)
	d.ValueTest()   // ValueTest(d)
	d.PointerTest() // PointerTest(&d)
	p.ValueTest()   // ValueTest(*p)
	p.PointerTest() // PointerTest(p)

}

/**
结果输出
Data: 0xc420014088
Value: 0xc420014098
Pointer: 0xc420014088
Value: 0xc4200140a0
Pointer: 0xc420014088
*/
