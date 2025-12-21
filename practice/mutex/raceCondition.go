package main

import (
	"fmt"
	"sync"
)

// RaceCondition: It Mean Multiple processes are trying to access or modify single resource

type Balance struct {
	amount int
	mu     sync.Mutex
}

func (amt *Balance) Add(wg *sync.WaitGroup) {
	defer func() {
		amt.mu.Unlock()
		wg.Done()
	}()
	amt.mu.Lock()
	amt.amount += 1
}
func main() {
	var wg sync.WaitGroup

	myBalance := Balance{
		amount: 0,
	}
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go myBalance.Add(&wg)
	}
	wg.Wait()
	fmt.Println("Total Value is:", myBalance.amount)
}
