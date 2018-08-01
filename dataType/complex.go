package main

import (
	"math/cmplx"
	"unsafe"
	"fmt"
)

func complex(){
	var a complex64 = 3.4;
	var b complex128 = 3 + 4i;
	fmt.Println( a)  // (3.4+0i)
	fmt.Println( cmplx.Abs( b)) // 5
	fmt.Println( unsafe.Sizeof(a))  // 8
	fmt.Println( unsafe.Sizeof(b))  // 16
}