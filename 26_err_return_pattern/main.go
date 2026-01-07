package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	// go don't use exceptions for normal failures
	// functions -> return errors as normal return value

	// val, err := something()
	// if err != nil {handle the error}

	if err := run(); err != nil {
		log.Fatal("Error wlile conversion:", err)
	}

}

func run() error {
	input := "3"
	level, err := parceLevel(input)

	if err != nil {
		return err
	}
	fmt.Println("selected level:", level)
	return nil
}

func parceLevel(s string) (int, error) {
	// (value, err)
	// nil error -> success
	// not nil -> failure

	// pattern
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("level must be a number")
	}

	if n < 1 || n > 5 {
		return 0, fmt.Errorf("level must between 1 and 5")
	}

	return n, nil
}
