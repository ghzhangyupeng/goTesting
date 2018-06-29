package main

import (
	"os"
	"encoding/gob"
	"runtime"
	"fmt"
)
// 定义文件路径
const file  = "./test.gob"

type User struct {
	Name,Pass string
}

func main()  {
	var datato = &User{"zhang", "pass"}
	var datafrom = new(User)
	// 把变量存储到文件中
	err := Save(file, datato)
	Check(err)
	err = load(file, datafrom)
	Check(err)
	fmt.Println(datafrom)
}

func Save(path string, object interface{}) error  {
	file,err := os.Create(path)
	if err == nil {
		fmt.Println(333)
		// 生成加密对象
		encoder := gob.NewEncoder(file)
		// 对数据进行加密
		encoder.Encode(object)
	}
	file.Close()
	return err
}

func load(path string, object interface{}) error  {
	file,err := os.Open(path)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(object)
	}
	file.Close()
	return err
}

func Check(e error)  {
	if e != nil {
		_,file,line,_ := runtime.Caller(1)
		fmt.Println(line,"--", file,"--",e)
		os.Exit(1)
	}
}





