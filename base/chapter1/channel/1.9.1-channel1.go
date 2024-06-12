package main

import "fmt"

func Hello(c chan string) {
  name := <- c
	fmt.Println("Hello", name)
}

func main() {
  ch := make(chan string)
  go Hello(ch)
  ch <- "World"
  close(ch)
  // fmt.Println("Hello World")
}
