package main

import (
	"fmt"
)

func createPerson(name string, age int8) *person{
    return &person{
		name: name,
		age: age,
	}
}

// go 允许这种返回局部变量的地址， c++不允许
func showPerson(){
	p1 := createPerson("lishuai", 22)
	
	fmt.Println(p1) // &{lishuai 22 <nil>}
}