package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println("Hello, 世界")
	fn := func(c rune) bool {
		return strings.ContainsRune(", |/", c)
	}

	res := strings.TrimFunc(" |/a,b|c/d", fn)
	// res = strings.TrimFunc(" |/abs sdg,/", fn)
	fmt.Println(res)
}
