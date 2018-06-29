package main

import (
	"time"
	"fmt"
)

func producer(ch chan int, num int)  {
	for {
		//time.Sleep(time.Millisecond)
		//if len(ch) != 0 {
		fmt.Println(<-ch, "----", num)
		if len(ch) == 0 {
			break
			//fmt.Println( "--====--", num)
		}
		//}
	}
	//close(ch)
}


func main()  {
	ch := make(chan int, 20)
	for i:=1; i<=20; i++ {
		//fmt.Println(i)
		time.Sleep(time.Millisecond)
		ch <- i
	}
	//time.Sleep(time.Second * 2)
	//for {
	//	//time.Sleep(time.Millisecond)
	//	if len(ch) != 0 {
	//		fmt.Println(<-ch, "----")
	//
	//	}
	//}
	for j:=1; j<=3; j++ {
		go func(ch chan int, num int) {
			producer(ch, num)
			fmt.Println("for is end")
		}(ch, j)
	}
	time.Sleep(time.Second * 2)
}
