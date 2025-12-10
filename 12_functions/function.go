package main

import "fmt"

// func add(num1, num2 int) int {
// 	return num1 + num2
// }

// func getLanguages()(string, string, bool) {
// 	return "golang", "javascript", true
// }

// func processIt(fn func(a int) int) {
// 	fn(2)
// }

// func processIt() func(a int) int {
// 	return func(a int) int {
// 		return 4
// 	}
// }

// Compute is a higher-order function: it takes a function op as argument
func compute(a int, b int, op func(int, int) int) int {
	return op(a, b)
}

// add is a function that matches signature: func(int, int) int
func add(x, y int) int {
	return x + y
}

// multipy is another function of same type
func multiply(x, y int) int {
	return  x * y
}

func main() {
	sum := compute(3, 4, add) // passed add as callback
	product := compute(3, 4, multiply) // passes multiply

	fmt.Println("Sum:", sum) // Sum: 7
	fmt.Println("Product:", product) // Product: 12

	// fn := func(a int) int {
	// 	return 2
	// }
	// fn := processIt()
	// fn(6)

	// var res = add(2, 4)
	// fmt.Println(res)
	// lang1, lang2, _ := getLanguages()
	// fmt.Println(lang1, lang2)
}
