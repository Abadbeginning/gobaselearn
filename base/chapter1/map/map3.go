package main

import (
	"fmt"
	"sort"
)

func main() {
	var arr map[int]int
	arr = make(map[int]int, 5)
	arr[0] = 33
	arr[1] = 11
	arr[2] = 22

	var b []int
	fmt.Println("排序前的值如下：")

	for k, v := range arr {
		fmt.Println(k, v)
		b = append(b, v)
	}

	sort.Ints(b)
	fmt.Println("排序后的值如下：")
	for k, v := range b {
		fmt.Println(k, v)
	}
	
}
