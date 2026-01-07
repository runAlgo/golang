package main

import "fmt"

func main() {
	// store a memory address of a value
	// &x -> address of x (makes a pointer)
	// *p -> dereference (go to that address and read/write)
	//It simply to change a value inside a function without returning it.

	score := 20
	fmt.Println("Before:",score)
	addScore(&score) // this line changed score into 20 to 25
	fmt.Println("After:", score)
}

func addScore(score *int) {
	*score = *score + 5
}



