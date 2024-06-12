package main

import "fmt"

func Func() func(string) string {
  a := "刷到过"
  return func(args string) string {
    return a + args
  }
}

func main() {
  f := Func()
  b := f("wdfgt")
  c := f("wdfgt")
  fmt.Println(b)
  fmt.Println(c)
}
