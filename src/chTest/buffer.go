package main

import (
	"bytes"
	"fmt"
)

func main()  {
	// buffer 是一个可以存储可变大小的字节缓存区，可以write和read方法操作，
	//buf := new(bytes.Buffer)
	var b bytes.Buffer
	b.Write([]byte("hello"))
	//arr := make(map[int]int, 5)
	//a,err := b.ReadByte()
	fmt.Println(b.String() )
}
