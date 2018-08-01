package main

import (
	"unsafe"
	"fmt"
)

func array(){
	var arr = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	var i int8
	for i = 0; i < 5; i++{
		fmt.Println(arr[i])
	}
	fmt.Println(unsafe.Sizeof(arr)) // 20
}