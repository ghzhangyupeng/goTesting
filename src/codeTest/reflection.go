package main

import (
	"math"
	"fmt"
)

type Square struct {
	side float32
}

type Circle struct {
	radius float32
}

type Shaper interface {
	Area() float32
}
func main()  {
	var  areaIntf Shaper
	// 生成一个地址类型（指针类型）
	sq1 := new(Square)
	fmt.Println(sq1,sq1.side)
	sq1.side = 5
	fmt.Println(sq1)
	//给接口赋值实现该接口的值
	areaIntf = sq1
	// 进行类型断言检测，看变量是否包含Square变量, 重点，括号里面，要是指针类型。
	if t,ok := areaIntf.(*Square);ok {
		fmt.Printf("the type of areaIntf is %T\n", t)
	}
	// 测试2 接口断言判断，也就是实现接口的变量，转化类型变量（*T），实现接口的变量必选是指针类型或者取地址类型。
	var area Shaper
	sq2 := Square{}
	sq2.side = 6
	area = &sq2
	if t,ok := area.(*Square);ok {
		fmt.Printf("the type of areaIntf is %T\n", t)
		fmt.Println(t)

	}
	// 进行对比，发现变量areaIntf是，没有Circle变量
	if u,ok := areaIntf.(*Circle); ok {
		fmt.Printf("the type of areaIntf is: %T\n", u)
	} else {
		fmt.Printf("areaIntf does not contain a variable of typeCircle")
	}

}

func (sq *Square)Area() float32 {
	return sq.side * sq.side
}

func (ci *Circle)Area() float32  {
	return ci.radius * ci.radius * math.Pi
}



