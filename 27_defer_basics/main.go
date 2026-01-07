package main

import (
	"errors"
	"fmt"
)

func main() {

	// fmt.Println("Case 1: success")
	// if err := task(true); err != nil {
	// 	fmt.Println("Error:", err)
	// }

	fmt.Println("Case 2: fail early")
	if err := task(false); err != nil {
		fmt.Println("Error:", err)
	}
}

func task(success bool) error {
	fmt.Println("start: resource aquired")

	defer fmt.Println("Cleanup: resource released")

	if !success {
		return errors.New("something went wrong. i am returning early.")
	}
	fmt.Println("work: doing something imp")
	fmt.Println("work: this work is done")
	return nil
}
