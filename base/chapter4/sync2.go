package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
  var mutex sync.Mutex
  wait := sync.WaitGroup{}
  var b = 0
  for i := 0; i < 10; i++ {
    wait.Add(1)
    go func(idx int) {
      mutex.Lock()
      b++
      fmt.Println("b --> ", b)
      mutex.Unlock()
      wait.Done()
    }(i)
  }
  time.Sleep(time.Second)
  wait.Wait()
}
