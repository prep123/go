package main

import "fmt"

func array(){
	var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	var i int8
	for i = 0; i < 5; i++{
		if balance[i] > 3 {
			fmt.Println(balance[i])
		}
	}
}