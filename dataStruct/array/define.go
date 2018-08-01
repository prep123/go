package main

import (
	"fmt"
)

func array(){
	var a1 [1]bool  // [false]
	var a2 = [2]int{1,4}
	var a3 = [...]float64{3.23, 325, 4543}
	var a4 [2][1]complex64  //  [[(0+0i)] [(0+0i)]] 两个长度为1的数组
	fmt.Println(a1, a2 ,a3, a4)
}