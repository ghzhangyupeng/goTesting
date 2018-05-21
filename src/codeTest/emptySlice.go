package main

import "fmt"

type sli struct {
	id int
	per int
	plat string
}

type SliType struct {
	sliArr []sli
}

func main()  {
	emptySlice := SliType{}
	fmt.Println(emptySlice, emptySlice.sliArr)
	for b,a := range emptySlice.sliArr{
		fmt.Println(b,"6666")
		fmt.Println("empty slice")
		if  a.id < 0{
			fmt.Println("sss")
		}
		if  a.id == 0{
			fmt.Println("0000")

		}
		fmt.Println("empty flag")
	}
	fmt.Println("empty flag")

	// range 循环 空数组 不报错 换句话说，就是不循环
	/*
	*	ss := []string{}  切片 普通初始化切片
	*	bb := make([]string, 3) make 出事切片，必须初始化长度
	*
	*
	*/
	ss := []string{}
	tt := make([]string, 4)
	for _, s := range ss {
		fmt.Println("empty slice flag")
		fmt.Println(s)
	}
	for k, b := range tt {
		fmt.Println("empty slice key:", k, "empty slice value:",b)
	}
	fmt.Println("a empty slice :", ss)
	fmt.Println("a empty slice :", tt)

}
