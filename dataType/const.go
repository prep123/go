package main

import (
	"fmt"
)
// C中常量是有类型的
// go 中 const 就像是文本替换
// 用到的时候决定类型
func constType(){
	const a = 3
	const b = "3"
	var p int = a * a
	var p1 complex128 = a + a
	var q string = b + b
	fmt.Println( p)   // 9
	fmt.Println( p1)  // (6+0i)
	fmt.Println( q)   // 33
}