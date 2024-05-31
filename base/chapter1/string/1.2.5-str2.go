package main

import (
	"fmt"
	"regexp"
	"strings" // Import the strings package
)

func main() {
  test := "i,love,code"
  str := test
  keywordSlice := strings.Split(str, ",")
  for _, v := range keywordSlice {
    reg := regexp.MustCompile("(?i)" + v)
    str = reg.ReplaceAllString(str, strings.ToUpper(v))
    // fmt.Println("val --> ", reg)
    fmt.Println(str)
  }

  fmt.Println("test --> ", test)
}
