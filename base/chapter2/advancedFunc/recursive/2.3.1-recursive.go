package main

import "fmt"

func main() {
  
  fmt.Printf("%d 是否偶数：%t\n", 7, Even(7))

  fmt.Printf("%d 是否奇数：%t\n", 2, Odd(2))

  fmt.Printf("%d 是否奇数：%t\n", 3, Odd(3))
}

func Even(n int) bool {
  if n == 0 {
    return true
  }
  return Odd(RecursiveSign(n) - 1)
}

func Odd(n int) bool {
  if n == 0 {
    return false
  }
  return Even(RecursiveSign(n) - 1)
}

func RecursiveSign(n int) int {
  if n < 0 {
    return -n
  }
  return n
}