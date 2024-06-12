package main

import "fmt"

func Func3() (r int) {
  defer func() {
    r += 3
  }()
  return 10  
}

func main() {
  fmt.Println(Func3())
}
