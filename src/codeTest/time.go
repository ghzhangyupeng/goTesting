package main

import (
	"time"
	"fmt"
)

func main()  {
	t := time.Now()
	uTime := t.Unix()

	fmt.Println(t, uTime)
}
