package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// jsonStr := `
	// 	{
	// 		"name": "zhangsan",
	// 		"age": 18
	// 	}	
	// `

	// var mapRes map[string]interface{}
	// err := json.Unmarshal([]byte(jsonStr), &mapRes)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(mapRes)

	instance := map[string]interface{}{
		"name": "zhangsan",
		"age":  18,
	}
	jsonStr, err := json.Marshal(instance)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonStr))
	
}
