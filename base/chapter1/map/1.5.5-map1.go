package main

import "fmt"

func main() {
	m := map[string]map[string]string{}
	m["a"] = map[string]string{}
	m["a"]["b"] = "c"
	m["height"] = map[string]string{}
	m["height"]["weight"] = "180"
	
	fmt.Println("二维 --> ", m, m["a"], m["a"]["b"])

	for key, value := range m {
		for k, v := range value {
			fmt.Println("k：", k, "v：", v)
		}

		fmt.Println("key：", key, "value：", value)
	}
}
