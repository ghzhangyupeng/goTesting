package main

import "fmt"

const (
	KBase = "BaseName"
)



type base struct {
	name string
}

func (b *base)get() string {
	return b.name
}

func Newbase(name string) *base  {
	return &base{name:name}
}

type per struct {
	*base
	name string
	zk string
}

//func (p *per)get()string  {
//	return  p.name
//}
func main() {
	person := &per{
		base:Newbase(KBase),
		name:"perName",
		zk:"zkstring",
	}

	// struct继承方式，A 继承 B
	/*
	*	如果是找属性，A如果没有，在去B中去相应的属性,如果B在没有，就报错
	*	如果是找方法，
	*	A和B中有相同的方法，当调用的时候，要根据func (p *base)get(),指定*base;
	*	A没有get()方法，B有get()方法，A调用get()方法，会调用B的get()的方法，忽略func (b *base)get(),指定的*base;
	*	A有get方法，B没有get()方法，A调用get()方法，正常调用
	*
	*/
	fmt.Println(person,person.base.name, person.get(), person.base.get())
}




