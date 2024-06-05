package main

import (
	"fmt"
	"sort"
)

func main() {
	
	a := []int{4, 3, 12, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 1, 5}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	fmt.Println(a)
	
}
