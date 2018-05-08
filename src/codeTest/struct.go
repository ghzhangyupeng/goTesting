package main

import (
	"fmt"
)

type test struct {
	name string
	aget int
}
type test1 struct {
	sex []string
	t []test
}

type test2 struct {
	t2 []test1
	num int
}

type test3 struct {

} 
func main()  {
	t := make([]test2, 3)
	t2 := test2{}
	t3 := test3{}
	t4 := test{}
	fmt.Println(t4,t, t2, t3)
	//j := json.Marshal(t3)
}
