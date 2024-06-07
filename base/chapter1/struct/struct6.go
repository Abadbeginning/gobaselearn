package main

import (
	"fmt"
	"sort"
)

type Programmer1 struct {
	Name string
	Age  int
	Lang string
}

type members []Programmer1

func (m members) Len() int {
	return len(m)
}

func (m members) Swap (i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m members) Less(i, j int) bool {
	return m[i].Age < m[j].Age
}

func main() {
	a := members{
		{"James", 27, "Go"},
		{"Jane", 22, "Java"},
		{"Jack", 29, "Python"},
	}
	sort.Stable(a)	
	fmt.Println(a)
}
