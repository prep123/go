package main

import (
    "fmt"
)

var arr = [...]int{1,2,5,43,734}
var s []int = arr[2:5]

func changeSlice(s []int){
    s[1] = 99999;
}

func testSlice(){
	fmt.Println( s) // [5 43 734]
	changeSlice( s)
	fmt.Println( arr) // [1 2 5 99999 734]
	fmt.Println( s) // [5 99999 734]
	arr[2] = 10000
	fmt.Println( s) // 10000 99999 734]
}