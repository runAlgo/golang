package main

import "fmt"

func fanInInt(channels ...chan int) chan int {
	out := make(chan int)

	for _, ch := range channels {
		go func(c chan int) {
			for v := range c {
				out <- v
			}
		}(ch)
	}

	return out
}


func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	merged := fanInInt(ch1, ch2, ch3)

	go func() {
		ch1 <- 10
		ch2 <- 20
		ch3 <- 30
	}()

	for i := 0; i < 3; i++ {
		fmt.Println("Received:", <-merged)
	}
}
