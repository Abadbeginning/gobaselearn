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
    field := getType.Field(i)
    fmt.Println("字段 --> ", field.Name, getValue.Field(i).Interface())
  }

  for i:=0;i<getType.NumMethod();i++{
    method := getType.Method(i)
    fmt.Println("方法 --> ", method.Name, method.Type, method.Type.NumIn(), method.Type.NumOut())
  }
}

func main() {
  pro := Programmer5{"刷到过", 12}
  
  GetFileAndMethod(pro)
}
