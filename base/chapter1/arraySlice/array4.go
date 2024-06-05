package main

import "fmt"

func main() {
	// fmt.Println("Hello, 世界")
	var array = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, -10}
	maxValue := array[0]
	maxIndex := 0
	for i := 0; i < len(array); i++ {
		if array[i] > maxValue {
			maxValue = array[i]
			maxIndex = i
		}
	}
	fmt.Println(maxValue, maxIndex)
}
