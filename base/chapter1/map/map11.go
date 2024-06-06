package main

import "fmt"

func main() {
	slice := make([]map[int]int, 3)

	for i := range slice {
		slice[i] = make(map[int]int, 6)
		slice[i][1] = 11
		slice[i][2] = 22
	}

	fmt.Println(slice)
}
