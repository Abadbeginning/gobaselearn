package main

import "fmt"

func main() {
  arr := Func2()
  fmt.Println(arr)
  arr[0]()
  arr[1]()
}

func Func2() []func() {
  b := make([]func(), 2, 2)
  for i:=0;i<2;i++ {
    b[i] = (func(j int) func() {
      return func() {
        fmt.Println(&j, j)
      }
    })(i)
  }
  return b
}
