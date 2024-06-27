package main

import "fmt"

func loop() {
  for i := 0; i < 20; i++ {
    fmt.Println("loop --> ", i)
  }
}

func main() {
  go loop()
  fmt.Println("最后 ---- ")
  loop()
}
