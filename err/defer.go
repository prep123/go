package main

import (
	"fmt"
)

// 不怕return ，defer 有自己的栈
func tryDefer(){
	defer fmt.Println( 1)
	defer fmt.Println( 2)
	fmt.Println( 3)
	return
    // fmt.Println( 4)
}
// 3
// 2
// 1