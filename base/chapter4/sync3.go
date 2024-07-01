package main

import "fmt"

func loop() {
  for i := 0; i < 20; i++ {
    fmt.Println("loop --> ", i)
  }
}

// Bug 修复：删除重复的 main 函数定义
func main() {
  go loop()
  fmt.Println("最后 ---- ")
  loop()
}
