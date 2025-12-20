package main

import (
	"fmt"
	"sync"
	"time"
)

// Rate-Limited API Worker Pool

type Job struct {
	ID int
}
// 1. Rate-Limited API Worker Pool
// Problem :
// I'm building a  backend service that processes 10,000 API requests, but:
// * Only 5 requests per second are allowed
// * Maximum 20 concurrent workers
// * Requests must not be dropped

// How would I design this using goroutines?

// Answer and Approach
// I am Using:
// * Worker pool (bounded concurrency)
// * Ticker for rate limiting
// * Buffered channel for job queue

func worker(id int, jobs <-chan Job, limiter <-chan time.Time, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		<-limiter // rate limit, It blocks this line for 200ms, then after next line runs
		fmt.Printf("Worker: %d processing job: %d\n", id, job.ID)
		// fmt.Println(100 * time.Millisecond)
		// time.Sleep(1000 * time.Millisecond)
	}
}

func main() {
	jobs := make(chan Job, 100)
	limiter := time.Tick(time.Second / 5) // time.Tick() returns a channel, the channel sends time.Time values, It sends one every duration "d"
	var wg sync.WaitGroup

	for i := 1; i <= 20; i++ {
		wg.Add(1)
		go worker(i, jobs, limiter, &wg)
	}

	for i := 1; i <= 100; i++ {
		jobs <- Job{ID: i}
	}
	close(jobs)

	wg.Wait()
}

// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// 1. Rate-Limited API Worker Pool
// Problem:
// I'm building a backend service that processes 10,000 API requests, but:
// * Only 5 requests per second are allowed
// * Maximum 20 concurrent workers
// * Requests must not be dropped

// How would I design this using goroutines?

// Answer and Approach
// I am Using:
// * Worker pool (bounded concurrency)
// * Ticker for rate Limiting
// * Buffered channel for job queue

// type Job struct {
// 	ID int
// }

// func worker(id int, w *sync.WaitGroup, jobs <- chan Job, limiter <- chan time.Time) {
// 	defer w.Done()
// 	for job := range jobs {
// 	<-limiter
// 	fmt.Printf("Worker: %d completed job: %d\n", id, job.ID)
// 	// time.Sleep(time.Second)
// 	}

// }

// func main() {
// 	var wg sync.WaitGroup
// 	jobs := make(chan Job, 100)

// 	limiter := time.Tick(time.Second / 5)

// 	for i := 1; i <= 20; i++ {
// 		wg.Add(1)
// 		go worker(i, &wg, jobs, limiter)
// 	}

// 	for i := 1; i <= 100; i++ {
// 		jobs <- Job{ID: i}
// 	}
// 	close(jobs)
// 	wg.Wait()
// }
