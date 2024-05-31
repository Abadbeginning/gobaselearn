package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "   cccHello, World!ccc     "
	str1 := strings.TrimSpace(str)
	str2 := strings.Trim(str, "c ")
	fmt.Println("str1 --> ", str1)
	fmt.Println("str2 --> ", str2)
}