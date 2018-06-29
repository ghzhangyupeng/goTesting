package main

import (
	"fmt"
)

func main()  {
	ch := make(chan int)
	go func(ch chan int) {
		fmt.Println("testing 1111")
		ch <- 33
		fmt.Println("testing 2222")

	}(ch)
	fmt.Println(<-ch)

	//time.Sleep(time.Second * 2)


}








