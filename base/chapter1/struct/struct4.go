package main

import (
	"fmt"
	"reflect"
)

type Programmer3 struct {
	// 标签
	Name string `json:"name" level:"12"`
	Age  int
	Lang string
}

func main() {
	// fmt.Println("Hello, 世界")
	var pro Programmer3 = Programmer3{}
	typeOfPro := reflect.TypeOf(pro)
	proFileName, ok := typeOfPro.FieldByName("Name")
	if ok {
		fmt.Println(proFileName.Tag.Get("json"), proFileName.Tag.Get("level"))
	}
}
