package main

import (
	"fmt"

	"github.com/runAlgo/go-modules/internal/greet"
)

func main() {
	name := "Kalu"
	msg1 := greet.Hello(name)
	fmt.Println(msg1)
}
