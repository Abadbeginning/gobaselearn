package main

import (
	"fmt"
	// "unicode/utf8"
)

func strTest() {
	str := "去年十月，as我们对外开源"
	// fmt.Println("长度 --> ", utf8.RuneCountInString(str))
	// fmt.Println("字节长度 --> ", len(str))
	// str1 := str[0:10]
	// fmt.Println("切片 --> ", str1)
	nameRune := []rune(str)
	// fmt.Println(len(nameRune))
	fmt.Println(string(nameRune[0:10]))
}

func Reversal(a string) (re string){
	b := []rune(a)
	for i := 0; i < len(b) / 2; i++ {
		b[i], b[len(b)-i-1] = b[len(b)-i-1], b[i]
	}
	// re = string(b)
	// return
	return string(b)
}

// 字符串翻转
func ReversalTest() {
	str := "123456789abc"
	strRev := Reversal(str)
	fmt.Println(str)
	fmt.Println(strRev)
}



func main() {
	// strTest()
	ReversalTest()
}
