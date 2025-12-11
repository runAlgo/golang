package main

import "fmt"

// func variadicFunc(nums ...interface{}) string {
// 	const numberMsg string = "This is Number"
// 	const stringMsg string = "This is String"

// 	for _, num := range nums {
// 		if reflect.TypeOf(num) == reflect.TypeOf(0) {
// 			return numberMsg
// 		}

// 		// Check for string
// 		if reflect.TypeOf(num) == reflect.TypeOf("") {
// 			return stringMsg
// 		}

// 	}

// 	return "Unknown type"

// }



// func sum(nums ...int) int {
// 	total := 0
// 	for _, num := range nums {
// 		total += num
// 	}
// 	return total
// }


// Another way to write variadic function
func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}
	return total
}
func main() {
	// fmt.Println(variadicFunc(1, 2, 3, 4, 5))
	// fmt.Println(variadicFunc("a", "b", "c"))
	

	// result := sum(1, 2, 3, 4)
	// fmt.Println(result)


	nums := []int{1, 2, 3, 4, 5}
	result := sum(nums...)
	fmt.Println(result)
}
