package main

import (
	"fmt"
	"reflect"
)

type Programmer5 struct {
	Name string
	Age  int
}

func (u Programmer5) ReflectCallFunc() {
  fmt.Println("阿呆发送给")
}

func GetFileAndMethod(input interface{}) {
  getType := reflect.TypeOf(input)
  fmt.Println("类型 --> ", getType.Name())
  getValue := reflect.ValueOf(input)
  fmt.Println("所有文件 --> ", getValue)

  for i:=0;i<getType.NumField();i++ {
    fmt.Println("字段 --> ", getType.Field(i).Name)
  }
}

func main() {
  
  fmt.Println("Hello World")
}
