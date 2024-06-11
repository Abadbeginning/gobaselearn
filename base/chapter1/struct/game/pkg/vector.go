package main

import "fmt"

type Vector struct {
  X float32
  Y float32
}

func (v Vector) Add(v2 Vector) Vector {
  return Vector{v.X + v2.X, v.Y + v2.Y}
}

func (v Vector) Sub(v2 Vector) Vector {
  return Vector{v.X - v2.X, v.Y - v2.Y}
}

func (v Vector) Mul(f float32) Vector {
  return Vector{v.X * f, v.Y * f}
}

func (v Vector) Div(f float32) Vector {
  return Vector{v.X / f, v.Y / f}
}

func main() {
  fmt.Println("Hello World")
}
