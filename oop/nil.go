package main

import (
	"fmt"
)

// nil 也能用
func nilStruct(){
	var p1 *person
	p1.hi()  // hello, world
}

func (p *person) hi(){
	fmt.Println("hello, world")
}