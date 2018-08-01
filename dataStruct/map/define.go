package main

import (
	"fmt"
)

func define(){
	dict := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}
	fmt.Println(dict["Red"])
}

func exist(){
	dict := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}
	fmt.Println(dict["OOO"])

	value, exists := dict["Blue"]
	if exists {
        fmt.Println(value)
    }
}

var dict = map[string]string{"Red": "#da1337", "Orange": "#e95a22"}
func change(dic map[string]string){
	dic["Red"] = "0000"
}

func prop(){
	change(dict)
	fmt.Println(dict) // map[Red:0000 Orange:#e95a22]
	// 值传递key，结构原来的value变了
}

