package main

import "fmt"

type User struct {
	Id    int
	Name  string
	Email string
	Age   int
}

func main() {
	user1 := User{
		Id:    1,
		Name:  "Kalu",
		Email: "kalu@email.com",
		Age:   23,
	}
	fmt.Println(user1)

	// struct fields are mutable by-default
	user1.Age = 24
	fmt.Println(user1)

	user2 := User{
		Email: "kamal@email.com",
		Name: "Kamal",
		Age: 24,
	}
	fmt.Println(user2)
}
