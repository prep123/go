package main

import (
	"fmt"
)

// 向切片尾部添加元素，如果没有超过容量，就是扩大切片长度
// 再修改底层数组
func appendSlice1(){
	var s = []int{1, 2, 3, 4, 5}
	s1 := s[1:3]
	s1 = append(s1,40)
    fmt.Println( s) // [1 2 3 40 5]
}


// 可以看到，很多时候数组是为了切片服务的
// 长度不够时，就会抛弃原来的数组，创建新的
func appendSlice2(){
	var s = []int{1, 2, 3}
	s1 := s[1:3]
	s1 = append(s1,40)

	fmt.Println( s) // [1 2 3]
	fmt.Println( s1) // [2 3 40]

	// 这时，s1的底层是新的数组了
	s1[0] = 20
    fmt.Println( s) // [1 2 3]
}