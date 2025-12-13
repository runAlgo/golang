package main

// 2. Generic Struct (Type Parameters)
// With Go 1.18+, Generics are now the standard for building
// reusable data structures. You need to practice 
// creating struct that can handle any type without losing type 
// safety (no more interface{} casting).

// ### What to Practice:
	// Creating "Wrapper" structs for API responses.
	// Building basic data structures (Stack, Queue) using [T any].

// Exercise: Build a generic APIResponse wrapper that works for any data model.

// T is a "Type Parameter". It can be User, Product, or anything else.
// type APIResponse[T any] struct {
// 	Status int `json:"status"`
// 	Message string `josn:"message"`
// 	data T		`json:"data"` // Data is now strictly typed as T
// }

// Usage
// type User struct { Name string}
// type Product struct {SKU string}


// func main() {
 // This response specifically holds a User
// userResp := APIResponse[User] {
// 	Status: 200,
// 	Data: User{Name: "Alice"},
// }
// }