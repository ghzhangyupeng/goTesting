package main

import "fmt"

func main()  {
	mapArr := make(map[string][]int, 3)

	if mapArr["1"] == nil {
		fmt.Println("map[][] is nil")
	} else {
		fmt.Println(mapArr)
	}
	// map testing

	m := mapArr["m"]
	fmt.Println(m, "----", mapArr, "---")
	str := make([][]int, 4)

	if  str[0] == nil{
		fmt.Println(" [][] is nil", str)
	} else {
		fmt.Println(" [][] is not nill", str)
	}
}
