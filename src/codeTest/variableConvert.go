package main

import (
	"fmt"
	"errors"
	"io"
)

type item struct {
	Name string
}

func (i *item) String() string {
	fmt.Println("sss")
	return fmt.Sprintf("item name:%v", i.Name)
}

type Persion struct {
	Name string
	sex string
}

//demo (p *Persion) String() string  {
//	return fmt.Sprintf("persion name:%v sex:%v", p.Name, p.sex)
//}

func Parse(i interface{}) interface{} {

	switch i.(type) {
		case string:
			return &item{
				Name:i.(string),
				}
		case []string:
			data := i.([]string)
			len := len(data)
			if len == 2{
				return &Persion{
					Name:data[0],
					sex:data[1],
				}
			} else {
				return nil
			}
		default:
			panic(errors.New("type match miss"))

	}

	return nil
}

type demo interface {
	name()
}
type test struct {
	d demo
}

func (n *test)name()  {
	fmt.Println("testing")
}

type demo2 struct{
	d2 demo
}
func main(){
	defer func() {
		if err := recover(); err != nil{
			fmt.Println(err)
		}
	}()
	d := test{}
	d.name()
	//fmt.Println(d.())

	//fmt.Println(name.(string))
	//it := item{"123"}
	//fmt.Println(it.String())
	//i1 := Parse("Apple").(*item)
	//fmt.Println(i1)
	//
	//p2 := Parse([]string{"zhang", "san"}).(*Persion)
	//
	//fmt.Println(p2)

	var x interface{} = 7  // x 的动态类型为int， 值为 7
	i := x.(int)           // i 的类型为 int， 值为 7
	fmt.Println(i)
	type I interface { m() }
	var y I
	//s := y.(string)        // 非法: string 没有实现接口 I (missing method m)
	r := y.(io.Reader)
	if r == nil {
		panic("ce")
	}

	fmt.Println(r)
}
