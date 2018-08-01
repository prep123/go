package main

import (
	"strconv"
	"time"
	"fmt"
)

func worker2(c chan int, id int){
	var i int
	for {
		i = <- c
		fmt.Println("Worker " + strconv.Itoa( id) + " received " + strconv.Itoa( i))
	}
}

func mutiWorker(){
	c := make(chan int)
	for i := 0; i < 100; i++ {
	    c <- i
	}
	
	for i := 0; i < 10; i++ {
	    go worker2( c, i)
	}

	time.Sleep( time.Microsecond)
}