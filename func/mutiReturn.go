package main

import (
	"fmt"
)

func muti() (a int, b error){
	return 1, nil
}

func mutiReturn(){
	a, err := muti()
	fmt.Println(a, err) // 1 <nil>
}