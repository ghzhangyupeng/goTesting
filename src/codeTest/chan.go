package main

import "fmt"

func Sum(values []int, resultChan chan int)  {
	sum := 0
	for _, value := range values{
		sum += value
	}
	resultChan <- sum
}

func main() {
	values := []int{1,2,3,4,5,6,7,8,9,10}
	val := values[1:8]
	fmt.Println(val)
	resultChan := make(chan int, 3)
	go Sum(values[:len(values)/2], resultChan)
	go Sum(values[len(values)/2:], resultChan)
	go Sum(values[len(values)/3:], resultChan)

	sum1,sum2,sum3 := <-resultChan, <-resultChan,<-resultChan
	fmt.Println(sum1,sum2,sum3)
}
