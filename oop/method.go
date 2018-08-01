package main

import (
	"fmt"
)

func (p person) showName() {
	fmt.Println(p.name)
}

func (p person) incAge(){
	p.age += 1
}

func (p *person) incAge2(){
	p.age +=1
} 

func method(){
	p1 := createPerson("lishuai", 22)

	p1.incAge()
	fmt.Println( p1.age)  // 22 值传递

	p1.incAge2()      // 自动传递指针
	fmt.Println(p1.age)   // 23 可以通过指针的方式达到间接引用传递
}