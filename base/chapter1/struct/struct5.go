package main

import "fmt"

type People struct {}

func (*People) SayHello(name string) {
	fmt.Println("Hello, ", name)
}

type Student struct {
	*People
}

func main() {
	name := "夏目"
	a := People{}
	a.SayHello(name)

	b := Student{&People{}}
	b.SayHello(name)
}
