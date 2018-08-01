package main

import (
	"unsafe"
	"fmt"
)

func float(){
	var a float32 = 3.4;
	var b float64 = 3;
	fmt.Println( unsafe.Sizeof(a))  // 4
	fmt.Println( unsafe.Sizeof(b))  // 8
}