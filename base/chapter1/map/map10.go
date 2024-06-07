package main

import (
	"fmt"
	"sort"
)

func main() {
	str := "aabbccdddffffffeeegjk"
	fmt.Println(str)
	f := frequencies(str)
	fmt.Println(f)
}

type frequency struct {
	char string
	fre int
}

func frequencies(s string) []frequency {
	m := make(map[string]int)
	for _, v := range s {
		m[string(v)]++
	}
	fmt.Println("m --> ", m)

	a := make([]frequency, 0, len(m))
	for c, f := range m {
		a = append(a, frequency{c, f})
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i].fre < a[j].fre
	})

	return a 
}