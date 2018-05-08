package main

import "fmt"

type demo interface{
	get()bool
	set()bool
}


type test struct {
	name string
	age int
}

func Newtest(t string,a int) *test {
	return &test{name:t,age:a}
}
func (t *test) get(s demo) string  {
	return t.name
}

func main()  {
	var t test
	t.name = "zhang"
	t.age = 12

	n := t.get(123)
	fmt.Println(n)
}
