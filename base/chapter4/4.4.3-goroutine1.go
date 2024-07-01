package main

import "fmt"

func generator(max int) <-chan int {
	out := make(chan int, 100)
	go func() {
		for i := 0; i < max; i++ {
			out <- i
		}
		close(out)
	}()
	fmt.Println("generator ", <-out)
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int, 100)
	go func() {
		for num := range in {
			out <- num * num
		}
		close(out)
	}()
	fmt.Println("square ", <-out)
	return out
}

func sum(in <-chan int) <-chan int {
	out := make(chan int, 100)
	// Sum := 0
	go func() {
		// var total
		var Sum int
		for v := range in {
			Sum += v
		}
		out <- Sum
		fmt.Println("sum ", <-out)
		close(out)
	}()
	return out
}

func main() {
	arr := generator(10)
	arr1 := square(arr)
	sum := <-sum(arr1)
	fmt.Println(sum)
}
