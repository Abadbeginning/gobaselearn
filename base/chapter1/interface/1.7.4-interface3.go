package main

import "fmt"

func main() {
  i := Minimum(1, 12, 3, 4, 5, 6, 7, 18, 91, 10).(int)
  j := Minimum(1.5, 12.2, 3.1, 4.124, 5.453, 6.0, 7.12, 18.124, 91.124, 10.12).(float64)
  fmt.Println(i)
  fmt.Println(j)
}

func Minimum(first interface{}, rest ...interface{}) interface{} {
  minimun := first

  for _, v := range rest {
    switch v.(type) {
      case int:
        if v := v.(int); v < minimun.(int) {
          minimun = v
        }
      case float64:
        if v := v.(float64); v < minimun.(float64) {
          minimun = v
        }
      case string:
        if v := v.(string); v < minimun.(string) {
          minimun = v
        }
    }
  }

  return minimun
}