package main

import "fmt"

func main() {
  arr := Func1()
  fmt.Println(arr)
  arr[0]()
  arr[1]()
}

func Func1() []func() {
  b := make([]func(), 2, 2)
  for i:=0;i<2;i++ {
    b[i] = func() {
      fmt.Println(&i, i)
    }
  }
  return b
}
