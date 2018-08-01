package main

import (
	"fmt"
	"math"
)

// go 中必须强制类型转换
func typeConvert(){
	var a, b int = 3, 4
	var c int
	c = int( math.Sqrt( float64( a * a + b * b)))
	fmt.Println(c)
}