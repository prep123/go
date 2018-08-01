package main

import (
	"fmt"
)

// 可以越切片的界，但是不能超过数组
func reslice(){
	arr := [...]int{1,2,3,4,5,6}
	s1 := arr[1:3]
	fmt.Println( len( s1))  // 2
	fmt.Println( cap( s1))  // 5
	s2 := s1[0:5]
	fmt.Println( s2) // [2 3 4 5 6]
}