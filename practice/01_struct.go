package main

// Anonymous Fields (Embedding)
// Go does not have classical inheritance, but it has embedding
// (sometimes called composition or an naonymous field). By embedding
// a struct (including it without a field name), the fields of the
// embedded struct are promoted to the outer struct.

type Engine struct {
	Power int // Horsepower
}

type Car struct {
	Name   string
	Engine // Embedded struct: Car now has access to the "Power" field directly
}

// type Product struct {
// 	id    string
// 	name  string
// 	price float32
// }

// func main() {
// myCar := Car{
// 	Name:   "Mustang",
// 	Engine: Engine{Power: 450},
// }
// Accessing the promoted field directly
// fmt.Println(myCar.Power) // Output: 450

// newProduct := Product{
// 	id:    "1",
// 	name:  "SunCream",
// 	price: 300.00,
// }

// fmt.Println(newProduct)
// }
