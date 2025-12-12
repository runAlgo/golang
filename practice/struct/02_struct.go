package main

// ## Struct and Pointers
// Structs are typically passed by value in Go. Using pointers is
// crucial for efficiency and when you need to modify the original
// struct.

// 1. Struct Pointers
// You can create a pointer to a struct using the & operator,
// or initialize it directly using new().

// type Person struct {
// 	Name string
// 	Age int
// }

// Accessing fields through a pointer
// go automatically dereferences the pointer (*p4).Name for you.

// func main() {
// Pointer to a struct
// p4 := &Person{Name: "David", Age: 50}
// 	fmt.Println(p4.Name, p4.Age)
// }
