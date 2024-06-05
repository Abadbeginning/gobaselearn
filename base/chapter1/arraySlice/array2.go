package main

import (
	"fmt"
	"reflect"
)

func main() {
	// fmt.Println("Hello, 世界")
	a := make([]int, 6)
	for i := 0; i < len(a); i++ {
		a[i] = i + 2
	}
	fmt.Println(arrayPostition(a, 6))
}

// 查找一个元素在数组中的位置
func arrayPostition(arr interface{}, d interface{}) int {
	array := reflect.ValueOf(arr)
	fmt.Println(array)
	for i := 0; i < array.Len(); i++ {
		v := array.Index(i).Interface()
		if v == d {
			return i
		}
	}
	return -1
}