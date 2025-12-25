package main

// import (
// 	"fmt"
// 	"sync"
// )

// // 1. ******* The Old Way: Communicating by Sharing Memory *******

// // Imagine you and a friend are working on a single notebook.
// // * The Shared Memory: The notebook.
// // * The Communication: To make sure you don't both write on the same
// // page the smae time and create a mess, you use a Lock(in programming, this called a Mutex).
// // * The Problem: You have to remember to look the book when you use it
// // and unlock it when you're done. If you forget, Your friend is stuck
// // writing forever (a Deadlock), or you both write at once and ruin the data (a Race Condition).

// // In this version, every goroutine (lightweight thread) tries to access the
// // same variable. We have to use a Mutex to "lock" the memory so only
// // one person touches it at a time.

// // type SafeCounter struct {
// // 	mu    sync.Mutex
// // 	value int
// // }

// // The focus: Protecting the variable.
// // The risk: If you forget c.mu.Unlock(), your program freezes.
// // func (c *SafeCounter) Increment() {
// // 	c.mu.Lock() // Lock the "notebook"
// // 	c.value++   // Change the shared data
// // 	fmt.Println(c.value)
// // 	c.mu.Unlock() // Unlock it for others
// // }

// // Value returns current value of the counter.
// // func (c *SafeCounter) Value() int {
// // 	c.mu.Lock()
// // 	// Lock is also needed for reading to ensure we don't read while
// // 	// another goroutinge is halfway through writing.
// // 	defer c.mu.Unlock()
// // 	return c.value
// // }

// // 2. ****** The Go Way: Sharing Memory by Communicating ****
// // Instead of sharing one notebook and fighting over a lock, imagine you
// // write your notes on a piece of paper and pass it to your friend.
// // * The Communication: Passing the paper through a 'pipe' (in Go, this is called a Channel).
// // * The Shared Memory: The data on that paper.
// // * The Benifit: Once you had paper off, you no longer have it. You can't accidentally overwrite it
// // because it's not in your hands anymore. Your friend now has full owenership of it.

// // In this version, no one "touches" the counter directly. Instead, we
// // send a message (communication) to a single worker who owns the data.

// func main() {
// 	// 1. ******* The Old Way: Communicating by Sharing Memory ******
// 	// counter := SafeCounter{}
// 	// var wg sync.WaitGroup

// 	// // Let's start 1000 "workers" (goroutines)
// 	// for i := 0; i < 1000; i++ {
// 	// 	wg.Add(1)
// 	// 	go func() {
// 	// 		defer wg.Done()
// 	// 		counter.Increment()
// 	// 	}()
// 	// }
// 	// // Wait for all 1000 increments to finish
// 	// wg.Wait()
// 	// // Print the final result
// 	// fmt.Println("Final Counter Value:", counter.Value())


// 	// 2. ****** The Go Way: Sharing Memory by Communicating ****
// 	// The "pipe" through which we communicate
// 	increments := make(chan int)
// 	done := make(chan bool) // bell

// 	// A WaitGroup to make sure we wait for all workers to send their signals
// 	var wg sync.WaitGroup
// 	// 1. The MANAGER (The "Owner" of the data)
// 	// This goroutine is the ONLY one that touches the "count" variable.
// 	// It shares the result by communicating, not by letting others in.
// 	go func() {
// 		count := 0
// 		for range increments {
// 			count++
// 		}
// 		// When the channel is closed, it prints the final result
// 		fmt.Println("Final Counter Value:", count)
// 		done <- true // task completed, ring the bell
// 	}()

// 	// 2. The WORKERS	
// 	// We start 1000 workers. They don't know 'count' exists.
// 	// They just know how to send a signal into the pipe.
// 	for i := 0; i < 1000; i++ {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			increments <- 1 // Send a signal
// 		}()
// 	}

// 	// 3. CLEANUP
// 	// Wait for all 1000 workers to finish sending
// 	wg.Wait()

// 	// Close the channel to tell the  Manager "no more data is coming"
// 	close(increments)
	
// 	// Wait for Manager
// 	<-done // stop until message comes from done, It stops the "main"
// 	// A small trick to let the Manager finish its final print before main exits
// 	fmt.Println("All signals sent!")
// }
