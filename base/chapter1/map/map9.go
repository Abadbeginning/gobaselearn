package main

import (
	"fmt"
	"sort"
)

func main() {
	counts := map[string]int{"a": 1, "b": 12, "c": 3}
	keys := make([]string, 0, len(counts))
	
	for key := range counts {
		keys = append(keys, key)
	}

	fmt.Println(keys)
	sort.Slice(keys, func(i, j int) bool {
		return counts[keys[i]] < counts[keys[j]]
	})

	for _, key := range keys {
		fmt.Println(key, counts[key])
	}

	fmt.Println(counts)
}
