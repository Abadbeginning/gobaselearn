package main

import (
	"fmt"
	"time"
)

func TimeLong(d time.Duration) <-chan struct{} {
  ch := make(chan struct{}, 1)
  go func() {
    time.Sleep(d)
    ch <- struct{}{}
  }()
  return ch
}

func main() {
  fmt.Println("Hello World")
  <-TimeLong(time.Second)
  fmt.Println("1s后", TimeLong(time.Second))
  <-TimeLong(time.Second)
  fmt.Println("再过1s后")
}
