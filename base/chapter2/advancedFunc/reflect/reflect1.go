package main

import (
	"fmt"
	"reflect"
)

func main() {
  var m float64 = 11.222
  fmt.Println(reflect.TypeOf(m))
  fmt.Println(reflect.ValueOf(m))
}
