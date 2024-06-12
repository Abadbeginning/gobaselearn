package main

import "fmt"

func main() {
  var str string
  var ptr *string
  var pptr **string

  str = "斯黛菲·格"

  ptr = &str
  pptr = &ptr
  fmt.Println("str ", str)
  fmt.Println(*ptr)
  fmt.Println(**pptr)
}
