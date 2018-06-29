package main

import "fmt"

type clusterStatus int8
const (
	STATUS_INIT   clusterStatus = iota
	STATUS_MASTER clusterStatus = iota
	STATUS_SLAVE  clusterStatus = iota
)
var num *int
func (c clusterStatus) String() string {
	switch c {
	case STATUS_INIT:
		return "init"
	case STATUS_MASTER:
		return "master"
	case STATUS_SLAVE:
		return "slave"
	}
	return ""
}
func main()  {
	fmt.Println(STATUS_INIT, STATUS_MASTER, STATUS_SLAVE)
	if num == nil {
		fmt.Println("空指针！！！")
	}
}


