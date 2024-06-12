package main

import "fmt"

func main() {
  f := func(x int) {
    fmt.Println(x)
  }
  f(12)
  
  // (func(x int) {
  //   fmt.Println(x)
  // })(223)
  
  // func(x int) {
  //   fmt.Println(x)
  // }(32)
}
