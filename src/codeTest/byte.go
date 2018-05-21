package main

import "fmt"

const ss  = iota

type Bidder interface {
	Bid(int) string
	test() string
}

type demo struct {
	name string
}

func (d *demo)test() string {
	return d.name+"333"
}

func (s *demo)Bid(i int) string  {
	return string(i)+"sss"
}
//Bidder接口，demo结构实现了Bidder所以可以作为返回值
func getBidder(name string) Bidder {
	return &demo{name:name}
}

func main() {
	ff :=  getBidder("zhang");
	fmt.Println(ff.test())
	var cbOkMsg = []byte("{\"retCode\":0}")
	var cbErrMsg = []byte("{\"retCode\":1}")

	fmt.Println(cbErrMsg, cbOkMsg)

}
