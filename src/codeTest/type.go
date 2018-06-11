package main

import (
	"fmt"
	"time"
	"os"
	"path"
)

type (
	Echo struct {
		ii int
		ss string
	}

	Route struct {
		Method string
	}
)

func main()  {
	e := &Echo{}
	os.Chdir(path.Dir(os.Args[0]))
	fmt.Println(e,os.Chdir(path.Dir(os.Args[0])))
	now := time.Now()
	h := uint32(now.Hour())
	weekday := now.Weekday()
	fmt.Println(weekday)
	weekday = (weekday + 6) % 7
	hourValue := 1 << h
	hourValue2 := int32(hourValue >> 1)
	fmt.Println(h, hourValue, hourValue2)
}

