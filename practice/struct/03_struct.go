package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) birthday() {
	p.Age++ // Modifies the original struct
}

func main() {
	// Methods with Pointer Receivers
	// when you define a method on a struct, using a pointer
	// receiver (*Person) is essential if the method needs to modify
	// the struct's underlying data.

	david := Person{
		Name: "Kalu",
		Age:  22,
	}
	david.birthday()

	fmt.Println(david.Age)
}