package main

import "fmt"

func main() {
  var ptr *string
  if ptr == nil {
    fmt.Println("ptr的值 ", ptr)
  }
}
