package main

import "fmt"

// In Go, the interface ensures that your code remains as flexible as the wall outlet.

// 1. The Interface (The Contract)
type Speaker interface {
	Speak() string
}

// 2. A Struct Implementation (A concrete object)
type Dog struct {
	Name string
}

// Dog implements the Speaker interface because it has Speak() method
func (d Dog) Speak() string {
	return "Woof!"
}

// 3. A Different Struct Implementaion
type Cat struct {
	Lives int
}

// Cat also implements the Speaker interface
func (c Cat) Speak() string {
	return "Meow."
}

// 4. Client Code uses the Interface
func Introduce(s Speaker) {
	// We don't care if 's' is a Dog or a Cat, only that it can Speak()
	fmt.Printf("I say: %s\n", s.Speak())
}


func main() {
	// Usage:
	dog := Dog{Name: "Buddy"}
	cat := Cat{Lives: 9}

	Introduce(dog) // Output: I say: Woof!
	Introduce(cat) // Output: I say: Meow!
}