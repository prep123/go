package main

import (
	"strconv"
	"time"
	"fmt"
)

func worker(c chan int){
	var a [10]int
	for i := 0; i < 10; i++{
		fmt.Println("excuting a[" + strconv.Itoa( i) + "]")
		a[i] = <- c
		fmt.Println( a[i])
	}
}

func defineChan(){
	c := make(chan int)
	go worker( c)
	c <- 0
	c <- 1
	c <- 2
    c <- 3

	time.Sleep( time.Microsecond)
}

// 可以看到。通道数据的读取是阻塞的，如果通道中没有数据，goroutine 就会暂停在哪里？
// 有点像读取键盘输入
/* 
excuting a[0]
0
excuting a[1]
1
excuting a[2]
2
excuting a[3]
 */