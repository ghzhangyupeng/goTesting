package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func walkFunc(path string, info os.FileInfo, err error) error {
	fmt.Printf("%s\n", path)
	return nil
}

func main() {
	//遍历打印所有的文件名
	filepath.Walk("/Users/zhangyupeng/web/adserverTest/service/dumpIndexPc", walkFunc)
}