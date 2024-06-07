package main

import "fmt"

type Progammer struct {
	Name string
	Age  int
	Lang string
}

func main() {
	// var members Progammer
	// members.Age = 12
	// members.Name = "James"
	// members.Lang = "Go"
	// fmt.Println(members)

	// members1 := Progammer{"James", 12, "Go"}
	// fmt.Println(members1)

	// var members2 = new(Progammer)
	// // (*members2).Age = 12 go做了处理
	// members2.Age = 12
	// members2.Name = "James"
	// (*members2).Lang = "Go"
	// fmt.Println(members2, *members2)

	var members3 = &Progammer{}
	members3.Age = 12
	members3.Name = "James"
	(*members3).Lang = "Go"
	fmt.Println(members3, *members3)
	// fmt.Println("Hello, 世界")
}
