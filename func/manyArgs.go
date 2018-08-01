package main

import "fmt"

func manyArgs(args ...int){
    for _, arg := range args {
		fmt.Println(arg)
	} 
}