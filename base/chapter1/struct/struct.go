package main

import (
	"fmt"
	"sort"
)

type Programmer struct {
	Name string
	Age  int
	Lang string
}

func main() {
	members := []Programmer{
		{"James", 27, "Go"},
		{"Jane", 22, "Java"},
		{"Jack", 29, "Python"},
	}
	
	fmt.Println("members --> ", members)
	sort.Slice(members, func(i, j int) bool {
		return members[i].Age < members[j].Age
		// return members[i].Age < members[j].Age || members[i].Name < members[j].Name
	})

	fmt.Println("members后 --> ", members)
}
