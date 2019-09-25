package main

import (
	"fmt"
	"strconv"
	"sync/atomic"
)

// 员工ID生成器 定义了一个函数类型
type EmployeeIdGenerator func(company string, department string, sn uint32) string

// 默认公司名称
var company = "Gophers"

// 序列号
var sn uint32

// 生成员工ID
func generateId(generator EmployeeIdGenerator, department string) (string, bool) {
	// 这是一条 if 语句，我们会在下一章讲解它。
	// 若员工ID生成器不可用，则无法生成员工ID，应直接返回。
	if generator == nil {
		return "", false
	}
	// 使用代码包 sync/atomic 中提供的原子操作函数可以保证并发安全。
	newSn := atomic.AddUint32(&sn, 1)
	return generator(company, department, newSn), true
}

// 字符串类型和数值类型不可直接拼接，所以提供这样一个函数作为辅助。
func appendSn(firstPart string, sn uint32) string {
	return firstPart + "-" + strconv.FormatUint(uint64(sn), 10)
}

func funcTest() {
	var generator EmployeeIdGenerator
	/*注意，函数generator是函数类型EmployeeIdGenerator的一个实现。实际上，
	只要一个函数的参数声明列表和结果声明列表中的数据类型的顺序和名称与某一个函数类型完全一致，前者就是后者的一个实现*/
	generator = func(company string, department string, sn uint32) string {
		return company + "-" + appendSn(department, sn)
	}

	fmt.Println(generateId(generator, "RD"))
}
