package main

import "fmt"

type HandlerCreator func(initParam *string) string
var name = make(map[string]string)
var mArr = make([]int,2)
func main(){
	name["n"] = "zhang"
	if name["m"] == "" {
		fmt.Println(name["m"], "--")
		fmt.Println("--")
	}

	if _, found := name["m"];found {
		fmt.Println("m", found)
	} else {
		fmt.Println("m", found)
	}

	if _, found := name["n"]; found {
		fmt.Println("n", found)
	} else {
		fmt.Println(name["m"])
	}




}
