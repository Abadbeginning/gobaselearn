package main

import "fmt"

func main() {
	// fmt.Println("Hello, 世界")
	test()
}

// 检查字符串是否在切片zhong 
func Exist(target string, array []string) bool {
	for _, element := range array {
		if element == target {
			return true
		}
	}
	return false
}

func test() {
	nameList := []string{"zhong", "li", "li", "zhang"}
	str1 := "zhong1"
	str2 := "c++"
	result := Exist(str1, nameList)
	fmt.Println(result)
	result = Exist(str2, nameList)
	fmt.Println(result)
	result = Exist("li", nameList)
	fmt.Println(result)
}