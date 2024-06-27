package main

import (
	"fmt"
	"time"
)

func main() {
  var b = 0
  for i := 0; i < 10; i++ {
    // go func(idx int) {
    //   b++
    //   fmt.Println("b --> ", b)
    // }(i)
    go func() {
      b++
      fmt.Println("b --> ", b)
    }()
  }
  
  time.Sleep(time.Second)
}
