package main

import "fmt"

func main() {
	varMap := make(map[int64]string)
	varMap[1] = "one"
	varMap[2] = "two"
	varMap[3] = "three"

	for _, num := range []int64{1, 2, 3, 4} {
		if _, ok := varMap[num]; ok {
			fmt.Println(num, "is in the map")
		} else {
			fmt.Println(num, "is not in the map")
		}
	}
}
