package main

import (
	"fmt"
	"sync"
)

func buy(n int) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for i := 0; i < n; i++ {
			ch <- fmt.Sprint("配件", i)
		}
	}()
	return ch
}

// 工序2组装
func build(in <-chan string) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for i := range in {
			ch <- "组装(" + i + ")"
		}
	}()
	return ch
}

func pack(in <-chan string) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for i := range in {
			ch <- "打包(" + i + ")"
		}
	}()
	return ch
}

// 加入扇出、扇入
func merge(ins ... <-chan string) <-chan string {
	var wg sync.WaitGroup
	out := make(chan string)

	p := func(in <-chan string) {
		defer wg.Done()
		for i := range in {
			out <- i
		}
	}

	wg.Add(len(ins))
	// fmt.Println("merge done", len(ins))
	for _, ch := range ins {
		go p(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	fmt.Println("merge done", <-out)
	return out
}

func main() {
	// // 工序1购买配件
	// parts := buy(5)
	// // 工序2组装
	// parts = build(parts)
	// // 工序3打包
	// parts = pack(parts)
	// for part := range parts {
	// 	fmt.Println(part)
	// }
	// 工序1购买配件
	parts := buy(7)
	// 工序2组装
	part1 := build(parts)
	part2 := build(parts)
	part3 := build(parts)
	computers := merge(part1, part2, part3)
	// 工序3打包
	parts = pack(computers)
	for part := range parts {
		fmt.Println(part)
	}
}
