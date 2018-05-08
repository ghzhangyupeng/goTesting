package main

import (
	"fmt"
	"reflect"
)

type demo struct {
	name string
}

func (d demo)getName() string  {
	return d.name
}
func Newdemo(name string) *demo {
	return &demo{name:name}
}
func main() {
	dd := Newdemo("ceshi")
	ceshi := reflect.ValueOf(Newdemo) 	// 这里参数是 函数类型，所以返回的函数地址
	c := ceshi.(&demo{})
	fmt.Println(c)
	demo := reflect.ValueOf(dd)
	demo3 := reflect.Indirect(ceshi)
	ss := "ss"
	d := []reflect.Value{reflect.ValueOf(ss)}
	demo2 := ceshi.Call(d)
	fmt.Println(dd, ceshi, demo, demo3, "--",ceshi.Kind(), demo2, d)


	var i int
	j := 3

	// reflect.ValueOf()的使用
	value := reflect.ValueOf(i) //使用ValueOf()获取到变量的Value对象
	value2 := reflect.ValueOf(j)
	fmt.Println(value, value2) // 输出 0 3


	// Value.Type() 和 Value.Kind()
	/* 相同点：都是获取对象和变量的类型，不同点：变量类型，返回的类型是一样的，差别在结构体对象。
	*/
	fmt.Println(value.Kind()) // 输出 int
	fmt.Println(value.Type()) //  输出 int

	type str struct {
		name string
	}
	var s str
	val := reflect.ValueOf(s)
	fmt.Println(val.Type()) // 输出 str
	fmt.Println(val.Kind()) // 输出 struct


	//获取变量的值 value.Interface() 该方法返回一个value的值，类型是interface
	var h int = 2

	//
	val1 := reflect.ValueOf(&h)
	val2 := reflect.Indirect(val1)
	val3 := val2.Interface()
	fmt.Println(val2.Kind(), val1.Kind(), val3)


}
