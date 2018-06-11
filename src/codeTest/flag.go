package main

import (
	"flag"
	"fmt"
)

var cluster string
var zk string

func main()  {
	flag.StringVar(&cluster, "cluster", "dev", "cluster name ( with prefix: bj | sh | hn | dev )")
	flag.StringVar(&zk, "zk", "", "zookeeper servers, seprated by comma")
	flag.Parse()
	fmt.Println(cluster)
}
