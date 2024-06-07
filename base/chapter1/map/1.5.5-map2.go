package main

import "fmt"

func main() {
	mainMap := map[string]interface{}{}
	subMap := make(map[string]string)
	subMap["a"] = "b"
	subMap["c"] = "d"
	mainMap["sub"] = subMap
	for key, val := range mainMap {
		fmt.Println("key:", key, "val:", val)
		// 接口先转换，在使用
		for subKey, subVal := range val.(map[string]string) {
			fmt.Println("subKey:", subKey, "subVal:", subVal)
		}
	}
}
