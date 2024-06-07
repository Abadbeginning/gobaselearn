package main

import (
	"encoding/json"
	"fmt"
)

type Item struct {
	Title string
	URL string
}

type Response struct {
	Data struct {
		Children []struct{
			Data Item
		}
	}
}

func main() {
	jsonStr := `{
		"data": {
			"children": [
				{
					"data": {
						"title": "Shirdon's Blog'",
						"url": "https://www.shirdon.com"
					}
				}
			]
		}
	}`

	res := Response{}
	json.Unmarshal([]byte(jsonStr), &res)
	fmt.Println(res)
}
