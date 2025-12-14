package main

import (
	"fmt"
)
func printSlice[T comparable, V string](items []T, name V) {
	fmt.Println(name)
	for _, item := range items {
	fmt.Println(item)
	}
}

// func printSlice[T int | string | bool](items []T) {
// 	for _, item := range items {
// 	fmt.Println(item)
// 	}
// }

// LIFO
// type Stack [T any] struct {
// 	elements []T
// }

func main() {
	// items := []int{1, 2, 3, 4}
	// names := []string{"Code","kalu", "Bablu"}
	// nums := []int{1, 2, 3, 4}
	boolean := []bool{true, false, false}
	printSlice(boolean, "kalu")

	// storeElem := Stack[int]{
	// 	elements: []int{1, 2, 3, 4},
		// elements: []string{"English", "DSA", "AI", "Robotics"},
// 
	// }

	// fmt.Println(storeElem)
}
