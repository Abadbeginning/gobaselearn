package main

import "fmt"

func main() {
	m := map[string]string {
		"a": "奥斯丁",
		"b": "巴西",
	}

	var sli []*string
	for k, v := range m {
		fmt.Printf("k: [%p] v:[%p], \n", &k, &v)

		sli = append(sli, &v)
		
	}

	for _, b := range sli {
		fmt.Println(*b)
	}
}
