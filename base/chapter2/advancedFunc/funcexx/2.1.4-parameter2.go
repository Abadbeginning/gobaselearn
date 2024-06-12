package main

import "fmt"

func Func5(i func(int, int) int) int {
  fmt.Println(i)
  return i(1, 2)
}

func main() {
  f := func(a int, b int) int {
    return a + b
  }
  fmt.Println(f)
  fmt.Println(Func5(f))
}
