package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println("Hello, 世界")
	str := []string{"I", "say", "the", "csgoclient"}
	res := strings.Join(str, "-")
	fmt.Println(res)
}