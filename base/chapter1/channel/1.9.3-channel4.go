package main

import "fmt"

var counter = func(n int) chan<- chan<- int {
	ch := make(chan chan<- int)
	go func() {
		for request := range ch {
			if request == nil {
				n++
			} else {
				request <- n
			}
		}
	}()
	return ch
}(0)

func main() {
	increase100 := func(done chan<-struct{})
	
	fmt.Println("Hello World")
}
