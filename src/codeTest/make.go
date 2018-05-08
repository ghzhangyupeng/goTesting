package main

import (
	"fmt"
)

func modify(array []int)  {
	array[1] = 33
}

func modify2(array [4]int)  {
	array[2] = 44
}
type doubleArr [4][4]int

func modifyDoubleArr(arr doubleArr) int {
	arr[3][3] = 12
	num := arr[3][3]
	return num
}

func main()  {

	slice := make([]int, 4, 5) //生成slice变量
	mapstring := make(map[string]int) // 生成map变量
	channel := make(chan int) // 生成通道变量
	fmt.Println("slice 变量输出：",slice)
	fmt.Println("map 变量输出：", mapstring)
	fmt.Println("chan 通道变量输出：", channel)
	/*
	*	make()  只作用于创建slice、map、channel这三种，这三种类型本来就是引用类型。
	*
	*
	*/
	// 判断make() 是否是变量是否是引用类型
	Slice := make([]int, 3, 5)
	Slice2 := []int{}
	arr2 := [4]int{1,2,3,4}
	arr3 := [4]int{}
	dArr := doubleArr{}
	modify(slice)
	modify2(arr2)
	modify2(arr3)
	fmt.Println(modifyDoubleArr(dArr))
	fmt.Println("有长度为3和容量为5的slice：",Slice)
	fmt.Println("没有长度和容量的slice：", Slice2)
	fmt.Println("长度为4的数组并且有4个值的arr", arr2)
	fmt.Println("长度为4的空数组arr", arr3)

	/*
	*	array和slice的区别:
	*	定义方式不同：arr := [3]int{1,2,3} 或者 arr := [4]int{} 数组必须有固定长度
	*	而切片不用，slice := []int{} 这种方式生成的切片，不能定义长度和容量
	*	slice := make([]int, 3, 5) 这种方式生成的切片，可以定义长度和容量
	*	使用场景：目前slice使用的场景，比较多，除非知道已知数据的长度。
	*
	*	实验测试：
	*	函数调用数组方式：例如 demo(arr [4]int) bool
	*	函数调用slice方式：例如 demo(Slice []int) bool
	*	主要的在参数部分，参数类型，slice，不用指定长度， arr 指定长度
	*
	*	引申：为什么，type doubleArr [4][4]int 这样就省的在func(arr [4][4]int) bool
	*	转化写法：demo(arr doubleArr) bool
	*/


}
