package main

import (
	"fmt"
	"reflect"
)

type Gopher struct {
  Id int
  Name string
  Level int
}

func (g Gopher) ReflectFuncHasArgs(name string, age int) {
  fmt.Println("有参数 --> ", name, age, g.Name)
}

func (g Gopher) ReflectFuncNoArgs() {
  fmt.Println("无参")
}

func main() {
  gro := Gopher{1, "阿呆", 12}
  getVal := reflect.ValueOf(gro)
  methodValue := getVal.MethodByName("ReflectFuncHasArgs")
  fmt.Println("methodValue --> ", methodValue)
  args := []reflect.Value{reflect.ValueOf("阿呆1"), reflect.ValueOf(12)}
  methodValue.Call(args)

  methodValue = getVal.MethodByName("ReflectFuncNoArgs")
  args = make([]reflect.Value, 0)
  methodValue.Call(args)
}
