package main

import "fmt"

func process(val *int) {
	if val == nil {
		return // Safety first!
	}
	fmt.Println(*val)
}

func main() {
	val := 20
	ptr := &val
	process(ptr)

	//*** Pro Tip for 2026: The new Keyword ****
	// If you want to create a pointer to a "blank" integer without
	// declaring a variable first, you can use the new function:

	// anotherPtr := new(int) // Allocates memory for an int, sets it to 0, returns the address
	// process(anotherPtr) // Prints 0

}
