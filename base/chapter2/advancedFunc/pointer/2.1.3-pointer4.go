package main

import "fmt"

func main() {
  var i int = 11
  var j int = 22
  
  fmt.Println(i)
  fmt.Println(j)

  swap(&i, &j)
  fmt.Println(i)
  fmt.Println(j)
}

func swap(x *int, y *int) {
  var temp int = *x
  *x = *y
  *y = temp
}