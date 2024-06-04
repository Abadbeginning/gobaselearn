package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println("Hello, 世界")
	s := "sd_you_are_the_best"
	// res := strings.Split(s, "_")
	// for i, v := range res {
	// 	fmt.Println("split --> ", res[i], v)
	// }


	res1 := strings.SplitN(s, "_", 2)
	fmt.Println("res1 --> ", res1)
	for i, v := range res1 {
		fmt.Println("SplitN --> ", res1[i], v)
	}


	res2 := strings.SplitAfter(s, "_")
	for i, v := range res2 {
		fmt.Println("SplitAfter --> ", res2[i], v)
	}

	res3 := strings.SplitAfterN(s, "_", 3)
	for i, v := range res3 {
		fmt.Println("SplitAfterN --> ", res3[i], v)
	}
}
