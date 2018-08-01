package main

import (
	"fmt"
)

func makeSlice(){
	s1 := make([]string, 5)  // 如果只指定长度，那么切片的容量和长度相等。
	s2 := make([]int, 3, 5)  // 使用长度和容量声明整型切片
	// s3 := make([]int, 5, 3)  // 容量小于长度的切片会在编译时报错
	s4 := []string{"Red", "Blue", "Pink"} // 通过切片字面量来声明切片 cap 等于 len
	// s5 := s4[2:4]  //  slice bounds out of range 为什么会爆出这个错误，不是有 cap ?

	fmt.Println( s1)  // [    ]
	fmt.Println( s2)  // [0 0 0]
	fmt.Println( s4)  // [Red Blue Pink]

	fmt.Println( cap( s1)) // 5
	fmt.Println( cap( s2)) // 5
	fmt.Println( cap( s4)) // 3
}