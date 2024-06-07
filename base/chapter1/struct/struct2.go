package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	user := User{
		Name: "James",
		Age:  27,
	}

	data := struct {
		Title string
		Users  User
	} {
		"info",
		user,
	}
	data.Users.Name = "Jack"
	fmt.Println(user)
	fmt.Println(data)
}
