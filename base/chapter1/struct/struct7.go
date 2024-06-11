package main

import "fmt"

// 正方形
type Square struct {
	sideLen float32
}

// 三角形
type Triangle struct {
	Bottom float32
	Height  float32
}

func (t *Triangle) Area() float32 {
	return t.Bottom * t.Height / 2
}

type Shape interface {
	Area() float32
}

func (s *Square) Area() float32 {
 return s.sideLen * s.sideLen 
}

func main() {
  t := &Triangle{Bottom: 10, Height: 5}
  s := &Square{sideLen: 20}
  shapes := []Shape{t, s}
  for n, _ := range shapes {
    fmt.Println(n, shapes[n], shapes[n].Area())
  }
}
