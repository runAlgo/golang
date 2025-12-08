package main

import "fmt"

// range :- It is used to iterating over data structure
func main() {
// Iterate a slice through range 

	// nums := []int{3, 4, 5, 4, 5}

	// for i, num := range nums {
	// 		fmt.Println(i, num)
	// }


// Iterate a map through range
	// users := make(map[int]string)
	// users[1] = "Kalu"
	// users[2] = "Sushil"
	// users[3] = "Mahesh"
	// users[4] = "Dipak"

	// users := map[int]string{1:"kalu", 2:"Kapil", 3:"Lallu", 4:"Anish"}
	// for k, v := range users {
	// 	fmt.Println(k, v)
	// }


	// unicode code point rune
	// i => starting bye of rune
	for i, v := range "kalu" {
		fmt.Println(i, string(v))
	}
}