package main

import (
	"fmt"
	"reflect"
)

func main() {
  var m float64 = 11.222
  p := reflect.ValueOf(&m)
  v := reflect.ValueOf(m)
  convertP := p.Interface().(*float64)
  convertV := v.Interface().(float64)
  fmt.Println(convertP)
  fmt.Println(convertV)
}
