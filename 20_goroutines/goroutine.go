package main

import (
	"fmt"
	"sync"
)

// goroutine: This is a lightwait worker which can do multiple tasks at a time.
// real-life example:- In a restrurant a waiter who is serving
// food to one customer at a time, but goroutine means multiple waiters
// who serves foods to multiple customer at a time, which means customer
// can't wait for a food, multiple waiters can serves at a time to all customers.

func serveFood(id int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("Served food to customer:", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go serveFood(i, &wg)
	}
	wg.Wait()
}
