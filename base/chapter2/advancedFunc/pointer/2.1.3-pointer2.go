package main

import "fmt"

func main() {
  const MAX int = 3
  arr := []int{11, 22, 33}
  var i int
  var ptr [MAX]*int

  for i = 0; i < MAX; i++ {
    ptr[i] = &arr[i] 
  }

  for i = 0; i < MAX; i++ {
    fmt.Println(i, *ptr[i])
  }
  // *ptr[0] = 100
  // fmt.Println(arr)
  // fmt.Println(*ptr[0])
}
