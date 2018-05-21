package main

import "fmt"

/*
*
*  如果单个值返回，类型是引用，返回类型，return 后面可以省略,返回值名字必须相等
*  如果是多个值返回，返回值包括，引用和传值类型，返回引用类型返回值，一定要相等，而且顺序要相等。
*  传值类型，不用和返回值的.(引用类型，只有map类型，返回值的名称，要和返回的定义形参名称相等)，其他类型，p
*
*/

func test()(freqValues map[string]int32,dd map[string]string, ch chan string) {
	freqValues = make(map[string]int32)
	freqValues["freValue"] = 123
	//ii := 333
	ii := make(map[string]string)
	ii["string"] = "string"
	chanal := make(chan string)
	return freqValues,ii, chanal
}

func test2() (ch chan string,freqValues map[string]int32,dd map[string]string) {
	freqValues = make(map[string]int32)
	//ii := 333
	ii := make(map[string]string)
	ii["string"] = "string"
	chanal := make(chan string)
	return chanal,freqValues,ii
}

func demo()(ii int)  {
	tt := 33
	return tt
}


func main() {

	ss, t, ch := test()
	fmt.Println(ss, t,ch)

	ee := demo()
	fmt.Println(ee)
}
