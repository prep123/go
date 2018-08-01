package main

import (
	"fmt"
)

var s = "hello 李帅"

// ，字符串的内容不能在初始化后被修改
func string(){
	/* s[0] = "H" */
	fmt.Println( s)
}

/* for i, ch := range s {
	fmt.Println(i, ch)//ch的类型为rune
}  */