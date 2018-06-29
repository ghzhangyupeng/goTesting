package main

import (
	"fmt"
	"time"
)

var c1 chan string = make(chan string)

func main()  {
	/*
	* 	这种测试只针对，没有缓存的通道的测试
	*/
	// 测试1、 错误写法 ：这种写法，会报错， 原因是，代码顺序执行，在往无缓存长度的通道写数据前， 必须有一个提前读取通道的操作
	//
	//func(){
	//	time.Sleep(time.Second * 2)
	//	c1 <- "result 1"
	//
	//}()
	//fmt.Println("c1 is", <-c1)

	//	测试2、正确写法：启动一个协程用于往通道写数据，在main中（主协程），读取通道数据。
	go func() {
		fmt.Println("chan start")
		c1 <- "result 1"
		fmt.Println("chan end")
	}()
	//c1 <- "result 2"
	//fmt.Println("c1 is", <-c1)

	// 测试3、同级别协程，一个协程写入数据，一个协程读取数据，不会报错。
	go func() {
		fmt.Println("c1 is", <-c1)
	}()

	time.Sleep(time.Second * 2)

	/*
	*	先决条件：在无缓冲信道条件下，协程之间，通信关系。
	*	1、在主协程中，有接收端，子协程，没有发送端会报错。
	*	2、在子协程中有发送端，在其他协程没有接收端，不会报错。
	*	3、在同级别协程中，如果对一个通道写入数据，如果通道没有缓存长度，不会报错。
	*
	*	个人理解：main主协程，执行优先级大于子协程优先级，在编译的过程中，校验主协程有没有进行通道接收，如果没有
	*/



}






