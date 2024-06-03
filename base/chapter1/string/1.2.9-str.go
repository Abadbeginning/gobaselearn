package main

import (
	"fmt"
	"regexp"
)

// 匹配中文字符
// 汉字范围
var cnRegexp = regexp.MustCompile("^[\u4e00-\u9fa5]$")

func main() {
	// fmt.Println("Hello, 世界")
	str := "我说的csgoclient123"
	StrFilterGetChinese(&str)
	fmt.Println(str)
}

func StrFilterGetChinese(src *string) {
	strNew := ""
	for _, v := range *src {
		if cnRegexp.MatchString(string(v)) {
			strNew += string(v)
		}
	}

	*src = strNew
}
