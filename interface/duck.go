package main

import (
	"fmt"
)

type duck struct {
	name string
}

func (d duck)sing() {
	fmt.Println( d.name + " is sing!")
}