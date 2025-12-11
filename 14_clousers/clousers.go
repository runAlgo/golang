package main

import "fmt"

// Clouser: A function that carries its surrounding variables with it.
func increment() func() int {
	var counter int = 0
	return  func() int {
		counter += 1
		return counter
	}
}

func main() {
	add := increment()
	fmt.Println(add())
	fmt.Println(add())
	fmt.Println(add())
	fmt.Println(add())
}