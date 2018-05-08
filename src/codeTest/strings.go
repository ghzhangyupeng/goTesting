package main

import (
	"fmt"
	"strings"
)

func main()  {

	var s  = "sssss"
	var d  = "dddd"

	// 字符串链接 用 +
	ssdd := s+d
	fmt.Println(ssdd)

	// 字符串长度 len()
	num := len(ssdd)
	fmt.Println(num)

	// 获取字符串中的字符
	dd := ssdd[7]
	fmt.Println(dd,"---")
	fmt.Printf("this \"%s\"", dd)

	var str  = "stringTest"
	var str2  = "stringtest"

	// strings.EqualFold 判断字符串是否相等（不区分大小写） 返回 bool
	fmt.Println(strings.EqualFold(str,str2)) // 返回 true

	str3 := "str_string"
	prefix := "str_"
	prefix2 := "1str_"
	// strings.HasPrefix 判断该前缀是否是字符串的前缀 返回 bool
	fmt.Println(strings.HasPrefix(str3, prefix)) // true
	fmt.Println(strings.HasPrefix(str3, prefix2)) // false

	suffix := "ing"
	suffix2 := "ings"
	//string.HasSuffix 判断该后缀是否是字符串的后缀 返回bool
	fmt.Println(strings.HasSuffix(str3, suffix))  // 返回 true
	fmt.Println(strings.HasSuffix(str3, suffix2)) // 返回 false

	contains := "_"
	contains2 := "+"
	// strins.Contains 判断该字符串是否在指定的字符串中 返回bool
	fmt.Println(strings.Contains(str3, contains)) // 返回 true
	fmt.Println(strings.Contains(str3, contains2)) // 返回 false

	// strings.ToLower() 字符串转化为小写
	fmt.Println(strings.ToLower("STRING"))
	// 字符串转化为大写
	fmt.Println(strings.ToUpper("string"))

	// strings.Title() 字符串首字母转化成大写
	title := "this is a title"
	title2 := "thisisatitle"
	fmt.Println(strings.Title(title))
	fmt.Println(strings.Title(title2)) // Thisisatitle

	count := "count count count"
	//strings.Count() 判断sep 字符串片段是否在字符串中, 个数
	fmt.Println(strings.Count(count, "count"))

	indexString := "indexString"
	// strings.Index() 返回指定字符串子在字符串的位置
	fmt.Println(strings.Index(indexString, "index"))

	replace := "replace string"
	// strings.Replace() 返回 替换字符串 返回字符串
	fmt.Println(strings.Replace(replace, "string", "int", 1))

	trim := "123trim123"
	// strings.Trim() 返回 去掉指定字符串的字符串
	fmt.Println(strings.Trim(trim, "123"))

	trimSpace := " trimSpace "
	// strings.TrimSpace() 返回去掉字符串两边的空格的字符串
	fmt.Println(trimSpace)
	fmt.Println(strings.TrimSpace(trimSpace))

}
