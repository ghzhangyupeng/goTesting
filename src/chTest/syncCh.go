package main

import (
	"sync"
	"fmt"
)

func main()  {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		fmt.Println("Goroutine 1")
		wg.Done()
	}()

	go func() {
		fmt.Println("Gorouteine 2")
		//wg.Done()
	}()

	wg.Wait()
}






