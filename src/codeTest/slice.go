package main

import "fmt"

func main()  {
	// 定义一个数组， 数组是有指定长度
	var  myArray [6]int = [6]int{1,2,3,4,5,6}

	// 基于数组创建一个数组切片
	var mySlice  []int = myArray[0:3]
	fmt.Println("Elements of myArray ")
	for _, v := range myArray{
		fmt.Println(v, "")
	}

	fmt.Println("Elements of mySlice ")

	for _,v := range mySlice {
		fmt.Println(v)
	}

	// myslice[:] 基于数组的所有元素创建的切片

	// myslice[:5] 基于数组的前5个元素创建的切片

	// myslice[5:] 基于数组的后面5个元素创建的切片

	// make 直接创建元素个数是5个数组切片，初始值是0
	slice1 := make([]int, 5)
	fmt.Println(slice1) // 打印出，5个零

	//make 直接创意见元素个数是5，容量是10, 初始值是0
	slice2 := make([]int, 5, 10)
	fmt.Println(slice2) // 打印出， 5个零， 不会答应出容量
	// len() 输出切边的长度
	fmt.Println(len(slice2))  // 打印 5
	// cap() 输出切边的容量
	fmt.Println(cap(slice2)) // 打印 10

	// 向切片中，添加元素 append([]Type, ele) return new slice（而不是在原来的切片上添加元素）
	slice3 := append(slice2, 1,2)
	fmt.Println(slice3)

	// 向切片中，添加一个切片 append([]Type, []Type...) return new slice 注意 参数 添加的新切片后面要 + ...
	slice4 := append(slice2, slice3...)
	fmt.Println(slice4)

	// 切片复制 copy(dst []Type, src []Type) int return 复制多少个
	sl := []int{1,2,3,4,5,6}
	sl1 := []int{7,8,9}
	copy(sl, sl1)
	fmt.Println(sl) // [7 8 9 4 5 6]
	s := []int{1,2,3,4,5,6}
	s1 := []int{7,8,9}
	copy(s1, s)
	fmt.Println(s1,copy(s1, s)) // [1 2 3]
}
