package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	// fmt.Println("Hello, 世界")

	// int32
	fn := func(c rune) bool {
		fmt.Println(c, reflect.TypeOf(c))
		return strings.ContainsRune(", |/", c)
	}
	fmt.Println(strings.FieldsFunc("a,b|c/d", fn))
}
