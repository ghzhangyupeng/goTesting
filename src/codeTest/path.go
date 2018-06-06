package main

import (
	"fmt"
	"path"
)

func main()  {
	pathDir := "/Users/zhangyupeng/go/goTesting"
	relPathDir := "../go/goTesting"

	// path.IsAbs() 判断路径是否是相对路径
	fmt.Println("---- path.IsAbs(string) bool 测试：----")
	fmt.Println("绝对路径判断：/Users/zhangyupeng/go/goTesting", path.IsAbs(pathDir))
	fmt.Println("相对路径判断：../go/goTesting", path.IsAbs(relPathDir))

	fmt.Println("---- path.Split(string) string 测试：----")
	filePath := "/Users/zhangyupeng/go/goTesting/path.go"
	// path.Split() 返回路径的目录和文件, 如果文件名称，那就返回：''，文件; 斜线，最后一个
	dir, filename := path.Split(filePath)
	fmt.Println("返回路径的目录：", dir, "返回路径的文件名称：", filename)

	//例如：只有路径，/Users/zhangyupeng/go/goTesting/, 注意：如果最后一个斜线，如果没有，goTesting就是文件
	dir1,filename1 := path.Split("/Users/zhangyupeng/go/goTesting/")
	fmt.Println("返回路径的目录：", dir1, "返回路径的文件名称：", filename1)

	//例如：只有文件的路径，path.go
	dir2,filename2 := path.Split("path.go")
	fmt.Println("返回路径的目录：", dir2, "返回路径的文件名称：", filename2)

	// path.Join(elem ...string) 返回一个路径，结果是经过简化的，所有的空格字符串元素会被忽略。
	JoinfilePath := path.Join(relPathDir, "path.go")
	fmt.Println("返回路径的目录:", JoinfilePath)

	//例如：路径和路径，/Users/zhangyupeng/go/goTesting/和relPathDir
	JoinfilePath2 := path.Join("/Users/zhangyupeng/go/goTesting/", relPathDir)
	fmt.Println("返回路径的目录:", JoinfilePath2)

	// path.Dir() 返回路径 例如：/Users/zhangyupeng/go/goTesting/，不会以斜线结尾 返回的是：/Users/zhangyupeng/go/goTesting
	fmt.Println("---- path.Dir(string) string 测试：----")
	Dirdir := path.Dir("/Users/zhangyupeng/go/goTesting/")
	fmt.Println("返回路径的：", Dirdir)

	// 例如：路径是带文件的，/Users/zhangyupeng/go/goTesting/path.go 返回的是：/Users/zhangyupeng/go/goTesting
	Dirdir2 := path.Dir("/Users/zhangyupeng/go/goTesting/path.go")
	fmt.Println("返回路径的：", Dirdir2)

	// 例如：路径只有文件名称，path.go 返回：.
	Dirdir3 := path.Dir("path.go")
	fmt.Println("返回的路径：", Dirdir3)

	fmt.Println("---- path.Base(string) string 测试：----")
	// 返回路径最后一个元素,如果元素是''，返回'.', 如果路径只是一个斜线，会返回'/',注意：不是符号
	fmt.Println("最后一个元素：", path.Base("/Users/zhangyupeng/go/goTesting/"))
	fmt.Println("最后一个元素：", path.Base("/Users/zhangyupeng/go/goTesting"))
	fmt.Println("最后一个元素：", path.Base("/Users/zhangyupeng/go/goTesting/path.go"))

	fmt.Println("---- path.Ext(string) string 测试：----")
	// path.Ext(string) string 返回路径的文件的扩展名，以'.'号为准，如果没有点号，就返回空。后缀名，包括点号
	fmt.Println("返回路径文件的扩展名：", path.Ext("/Users/zhangyupeng/go/goTesting/path.go"))
	fmt.Println("返回路径文件的扩展名：", path.Ext("/Users/zhangyupeng/go/goTesting/"))

	fmt.Println("---- path.Ext(string) string 测试：----")
	// path.Clean(string) string 通过单纯的此番操作，返回和path代表同一地址的最短路径.
	fmt.Println("返回最短路径：", path.Clean("/Users/zhangyupeng/go/goTesting/path.go"))
	fmt.Println("返回最短路径：", path.Clean("/Users/zhangyupeng/../../go/goTesting/path.go"))

	fmt.Println("---- path.Match(string) string 测试：----")
	flag,err := path.Match("path.go","path.go")
	fmt.Println("返回最短路径：", flag, err)

}




