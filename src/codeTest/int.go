package main

import (
	"fmt"
	"strings"
)


const (
	keyTypeAll = iota
	keyTypeSum
)


func main()  {
	var testInt int
	fmt.Println(testInt)

	test := keyTypeSum
	fmt.Println(test)

	adgkeys := strings.Split("cpm#6#7900", "#")
	keys := strings.Split(adgkeys[2], "_")
	fmt.Println(adgkeys, adgkeys[2], keys)
}
