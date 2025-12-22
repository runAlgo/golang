package main

import (
	"fmt"
	"os"
)

// delete file

func main() {
	// f, err := os.Open("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()

	e := os.Remove("example2.txt")
	if e != nil {
		panic(e)
	}
	fmt.Println("delete file successfully.")
}