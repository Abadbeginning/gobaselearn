package main

import "fmt"

func main() {
	array := []int{1, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 1, 5, 7}
	res := removeDuplicates(array)
	fmt.Println(res)
}

func removeDuplicates(array []int) []int {
	if len(array) == 0 {
		return nil
	}

	left, right := 0, 1
	for ;right < len(array);right++ {
		if array[left] != array[right] {
			left++
			array[left] = array[right]
		}
	}
	return array[:left+1]
}