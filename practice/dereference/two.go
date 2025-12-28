package main

// 2. Automatic Dereferencing
// In modern Go, you rarely need to write (*p).Field. Go provides implicit
// dereferencing for selectors (struct fields) and methods to keep code clean.

// type User struct {
// 	Name string
// }

// func main() {
// 	u := &User{Name: "Alice"}

// 	// Explicit (Cumbersome)
// 	fmt.Println((*u).Name)

// 	u.Name = "Kalu"

// 	// Implicit/Automatic (Professional way)
// 	fmt.Println(&u.Name, u.Name)
// }
