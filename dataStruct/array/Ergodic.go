package main

import (
	"fmt"
)

func ergodic1(){
	var a = [...]float64{3.23, 325, 4543}
	for i := 0; i < len(a); i++ {
		fmt.Println(i, a[i])
	}
}

func ergodic2(){
	var a = [...]float64{3.23, 325, 4543}
	for i, v := range a {
		fmt.Println(i, v)
	}
}
// 0 3.23
// 1 325
// 2 4543