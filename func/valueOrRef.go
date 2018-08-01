package main

import (
	"fmt"
)

func changeArray(a [4]int8){
	a[1] = 127
	fmt.Println(a)  // [1 127 4 6]
}

func pointArray(a *[4]int8){
	// (*a)[1] = 127 Go 指针很单纯
	a[1] = 127
	fmt.Println(a)  // &[1 127 4 6]
}

func array(){
	var a = [...]int8{1,2,4,6}
	fmt.Println(a)  // [1 2 4 6]
	changeArray( a)  
	fmt.Println(a)  // [1 2 4 6]

	pointArray(&a)
	fmt.Println(a)  // [1 127 4 6]
}
