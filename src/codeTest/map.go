package main

import "fmt"

type Person struct {
	id int
	name string
}
type M map[string][]string
func get(m M, key string) (string)  {
	if m == nil {
		return ""
	}
	vs := m[key]
	if len(vs) == 0 {
		return ""
	}
	return vs[0]
}
func main()  {
	var per map[string]Person
	per = make(map[string]Person)

	// 往 map 插入元素
	per["1"] = Person{1, "zhang"}
	per["2"] = Person{2,"huo"}

	// 从 map 查找元素 例如查找元素 键值为2的
	// 返回两个值， 一个是返回结果，另一个是返回是否存在状态值(bool) ok true 存在，false 不存在
	per2, ok := per["2"]
	if  ok == true{
		fmt.Println("exist is ",per2)
	} else {
		fmt.Println("no exist")
	}
	// 如果不存在，返回一个空的类型值，状态值 false
	if per3, ok := per["4"]; ok == false {
	  fmt.Println(per3) // 打印 {0 }
	}

	// 从map中，删除元素
	delete(per, "2")
	fmt.Println(per)

	per["maptest"] = Person{3,"maptest"}
	for _,v := range per{
		fmt.Println("map for loop : ", v)
	}

	mm := M{"zhang":{"yupeng"}, "huo":{"yujie"}}

	name := get(mm, "huo")
	fmt.Println(name)



	m := make(map[string]string, 6)
	m["name"] = "ceshi"
	m["ceshi"] = "ceshi44"
	for xx, nn := range m  {
		fmt.Println(xx, nn)
	}
}
