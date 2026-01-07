package main

import "fmt"

type User struct {
	Name string
	Age  int
}


func main() {
	u := User{
		Name: "Sajan",
		Age:  22,
	}
	fmt.Println("Before:", u.Age)
	u.Birthday()
	fmt.Println("After:", u.Age)
}

// Methods pointer receiver
func (u *User) Birthday() {
	u.Age++
}
