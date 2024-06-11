package main

import (
	"fmt"
	"person"
)

func main() {
  s := new(person.Student)
  s.setScore(100)
  fmt.Println(s.getScore())
}
