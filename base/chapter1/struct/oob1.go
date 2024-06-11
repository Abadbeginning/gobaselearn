package main

import "fmt"

type Circle struct {
  Radius float32
}

func (c *Circle) Area() float32 {
  return 3.14 * c.Radius * c.Radius
}

func main() {
  r := Circle{Radius: 10}
  fmt.Println(r.Area())
}
