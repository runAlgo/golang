// Generics:- it's simple logic is write one pice of code that works for many types,
// without repeating yourself or losing type safety. They're most useful
// when the logic is the same but the types differ.

// Real-life Use Case #1 --- Reusable Utility Functions
// Example: A Map function
// Imaging you want to transform a list of things -- like doubling
// numbers or making all names uppercase.

// Before generics you might write separate function for []int and
// []string, but with generics:

package main

import (
	"fmt"
)

// Map now uses generics correctly:
// T is the input element type, U is the output element type.
func Map[T any, U any](items []T, fn func(T) U) []U {
	result := make([]U, len(items))
	for i, v := range items {
		fmt.Println("Befor executing fn inside range loop:", v)
		result[i] = fn(v)
		fmt.Println("After executing fn inside range loop:", result[i])
	}
	return result
}

// func main() {
// 	names := []string{"bob", "sum"}

// 	// Call Map with names slice + a function to transform each string
// 	upper := Map(names, func(s string) string {
// 		return strings.ToUpper(s)
// 	})

// 	// Print the result
// 	fmt.Println(upper) // Output: [BOB SUM]
// }
