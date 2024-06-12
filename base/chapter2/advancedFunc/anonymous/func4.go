package main

import "fmt"

func returnFunc(a int) (func(int) int, func(int) int) {
  return func(b int) int {
    return a + b
  }, func(b int) int {
    return a - b
  }
}

func main() {
  addFunc, subFunc := returnFunc(10)
  fmt.Println(addFunc(5))  // Output: 15
  fmt.Println(subFunc(5))  // Output: 5

}
