package main

import (
	"strconv"
	"time"
	"fmt"
)

func worker3(c chan int, id int){
	var i int
	for {
		i = <- c
		fmt.Println("Worker " + strconv.Itoa( id) + " received " + strconv.Itoa( i))
	}
}

func mutiChannel(){
	var channels [10]chan int
	for i := 0; i < 10; i++ {
	    for j := 0; j < 20; j++ {
			channels[i] <- j
		}
	}
	fmt.Println(channels[3])
	
	/* for i := 0; i < 10; i++ {
	    go worker3( channels[i], i)
	} */

	time.Sleep( time.Microsecond)
}