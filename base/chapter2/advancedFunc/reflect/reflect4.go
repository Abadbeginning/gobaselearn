package main

import (
	"fmt"
	"reflect"
)

func main() {
  var money float64 = 66.11
  fmt.Println("指针 --> ", money)
  p := reflect.ValueOf(&money)
  v := p.Elem()
  fmt.Println("指针类型 --> ", v.Type())
  fmt.Println("指针值 --> ", v.CanSet())

  v.SetFloat(123.33)
  fmt.Println("新值 --> ", money)

  
}
