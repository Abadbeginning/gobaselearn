package main

import "fmt"

func main() {
  var str string = "hello"
  var name *string = &str
  fmt.Println("str ", &str)
  fmt.Println("name地址 ", name)
  fmt.Println(*name)
}
