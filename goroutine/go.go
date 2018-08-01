package main

import (
	"time"
	"fmt"
	"runtime"
)

func testGo(count int){
	for i := 0; i < count; i++ {
		go func(i int){
			for{
				fmt.Printf("hello %d! \n", i)
			}
		}(i)
	}
	time.Sleep( time.Millisecond)
}

func incGo(){
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int){
			for{
				a[i]++
				runtime.Gosched() // 不加这句话的话，一个协程霸占cpu，结束命令无法执行
			}
		}(i)
	}
	// time.Sleep( time.Millisecond)
	// fmt.Println( a)  // [536 517 464 441 473 498 457 429 425 430]

	time.Sleep(time.Minute)  // goCPU.git   超级系统占用，不超过四个线程
}


func order(){
	go testOrder()
	fmt.Println( 9)
}
// 只输出9， go语句还没有来得及执行主程序就挂了
func testOrder(){
	fmt.Println( 1)
}

