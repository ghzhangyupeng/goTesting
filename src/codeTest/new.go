package main

import (
	"fmt"
	"unsafe"
)
type SyncedBuffer struct {
	str string
	num int
	sum int
}
func main()  {
	var p  *[]int = new([]int)
	fmt.Println(p)

	m := new(SyncedBuffer)
	m1 := SyncedBuffer{}
	c := &m1
	m2 := &SyncedBuffer{}
	m3 := &SyncedBuffer{"123",123, 44}
	m4 := SyncedBuffer{"123", 123, 33}
	m3.str = "m3Test"
	fmt.Println("new m 变量输出：", m)
	fmt.Println("普通赋值 m1 变量输出：",m1)
	fmt.Println("&Type赋值 m2 变量输出：", m2)
	fmt.Println("&Type赋值 带有初始值的m3 变量输出：", m3)
	fmt.Println("&Variable赋值方式 c 变量输出：", c)

	fmt.Println("new m 变量占用内存情况：",unsafe.Sizeof(m))
	fmt.Println("普通赋值 m1 变量占用内存情况：",unsafe.Sizeof(m1))
	fmt.Println("&Variable赋值方式 c 变量占用内存情况：", unsafe.Sizeof(c))
	fmt.Println("&Type赋值 m2 变量占用内存情况：", unsafe.Sizeof(m2))
	fmt.Println("&Type赋值 带有初始值的m3 变量占用内存情况", unsafe.Sizeof(m3))
	fmt.Println("&Variable赋值方式 m4 变量占用内存情况：", unsafe.Sizeof(m4))

	/*
	*	new() 用于返回类型的指针，这个操作：分两步，第一步：初始化一个变量为零的值，第二步：生成一个指向类型变量的地址。
	*	&Type 返回类型变量的地址，这个操作：分两步，第一步：初始化一个变量（变量值可以不为零）的值，第二步：取出类型变量的地址并返回。
	*	通过实验测试：
	*	new()的作用和&Type作用是一样的，都是返回类型变量的地址(指向类型变量的指针)
	*	相同点：都是返回指向变量的地址，
	*	不同点：new()不支持初始化变量的值(默认值都是零)，&Type支持初始化自定义变量的值
	*	使用场景：
	*	new() 变量类型里面的值，是零值，就可以，不需要动态变化，比如type struct 属性全是锁
	*	&type 可以自定义初始化变量的值，也可以像new()一样，默认生成值(零值)，这样做的好处，不用想c语言那样，申请一个变量，在取地址，
	*	赋值给指针变量，简化操作，注重实用。
	*
	 */

}
