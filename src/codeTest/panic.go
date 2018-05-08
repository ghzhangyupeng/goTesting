package main

import (
	"strings"
	"fmt"
)

func equal()  {
	param := "123"
	param2 := "2"
	flag := strings.EqualFold(param, param2)
	if !flag {
		panic("two string is not equal")
	}
}

func main()  {

	defer func() {
		fmt.Println("defer 2")
	}()
	defer func() {
		fmt.Println("defer 1")
	}()

	defer func() {
		if err := recover(); err != nil{
			fmt.Println(err, "defer 0")
		}
		fmt.Println("testing running 1")
	}()

	panic("fault")

	fmt.Println("testing running 2")
	/*
	* 	defer 的执行顺序 是类似冒泡方式，从底往上执行 例如：defer 0 ，defer 1，defer 2
	* 	panic() 抛出错误，从此之后的代码， 不再执行 例如： testing running 2 不会被打印出来
	*	recover() 用于接收panic() 抛出的错误， 从此之后的代码，继续执行 例如：testing running 1 将会被打印出来
	*			  recover()必须在函数前面，否则后面panic()抛出异常，是捕获不到的。
	*/
}
