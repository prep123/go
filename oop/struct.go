package main

import (
	"fmt"
)

type person struct {
	name string
	age int8
	father *person
}

func showStruct(){
	var p1 = person{
		name: "lishuai",
		age: 22,
	}
	fmt.Println(p1) // {lishuai 22 <nil>}
}

