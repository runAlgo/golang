package main

import "fmt"

// for -> In go for is the ony concept for looping, Inside go while loop is not available. We can loop through for loop.
// for -> only construct in go for looping
func main() {
	// while loop
	// i := 1
	// for i <= 3 {
	// 	fmt.Print(i)
	// 	i = i + 1
	// }

	// infinite loop
	// for {
	// 	fmt.Println(1)
	// }

	// classic for loop

	// for i := 0; i <= 3; i++ {

	// 	if i == 1 {
	// 		break
	// 		// continue
	// 	}
	// 	fmt.Println(i)
	// }


	

	// 1.22 go version (range)

	for i := range 5 {
		fmt.Println(i)
	}
}
