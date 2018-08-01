package main

import (
	"fmt"
)

// 支持高阶函数
func adder() func (value int) int {
	sum := 0
	return func (value int) int {
		sum += value
		return sum
	}
}

// 支持闭包
func add(){
	adder := adder()
	for i := 0; i < 10; i++ {
	    fmt.Println( adder( i))	
	}
}