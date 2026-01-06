// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

//***** Fan out / Fan in ****
// type Result struct {
// 	Value string
// 	Err   error
// }

// func worker(url string, wg *sync.WaitGroup, resultChan chan Result) {
// 	defer wg.Done()
// 	time.Sleep(time.Microsecond * 50)

// 	fmt.Printf("image processed: %s\n", url)
// 	resultChan <- Result{
// 		Value: url,
// 		Err:   nil,
// 	}
// }

// func main() {
// 	fmt.Println("Welcome to go concurrency.")

// 	var wg sync.WaitGroup
// 	urls := 10

// 	resultChan := make(chan Result, urls)

// 	startTime := time.Now()
// 	// start workers (producers)
// 	wg.Add(urls)
// 	for i := 0; i < urls; i++ {
// 		go worker(fmt.Sprintf("image_url%v.png", i), &wg, resultChan)
// 	}

// 	go func(){
// 	wg.Wait()
// 	close(resultChan)
// 	}()

// // consume results (consumer)
// 	for result := range resultChan {
// 		if result.Err != nil {
// 			fmt.Println("error:", result.Err)
// 			continue
// 		}
// 		fmt.Println("received:", result.Value)
// 	}

// 	fmt.Printf("it took %v µs.\n", time.Since(startTime).Microseconds())
// 	fmt.Println("all done✅")
// }

// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

/// **** WORKER POOL Pattern ******

// func worker(jobsChan chan string, wg *sync.WaitGroup, resultchan chan Result) {
// 	defer wg.Done()

// 	for job := range jobsChan {
// 		time.Sleep(time.Microsecond * 50)
// 		// fmt.Printf("image processed: %s\n", job)

// 		resultchan <- Result{
// 			Value: job,
// 			Err:   nil,
// 		}
// 	}
// 	fmt.Printf("Worker shutting down\n")
// }

// type Result struct {
// 	Value string
// 	Err   error
// }

// func main() {
// 	jobs := []string{
// 		"image_1.png",
// 		"image_2.png",
// 		"iamge_3.png",
// 		"image_4.png",
// 		"image_5.png",
// 		"image_6.png",
// 		"image_7.png",
// 		"image_8.png",
// 		"image_9.png",
// 		"image_10.png",
// 	}
// 	var wg sync.WaitGroup
// 	titalWorkers := 5
// 	resultChan := make(chan Result, 50)
// 	jobsChan := make(chan string, len(jobs))

// 	startTime := time.Now()

// 	for i := 1; i <= titalWorkers; i++ {
// 		wg.Add(1)
// 		go worker(jobsChan, &wg, resultChan)
// 	}

// 	go func() {
// 		wg.Wait()
// 		close(resultChan)
// 	}()

// 	// send the jobs
// 	for i := 0; i < len(jobs); i++ {
// 		jobsChan <- jobs[i]
// 	}

// 	close(jobsChan)

// 	for result := range resultChan {
// 		fmt.Printf("✅ Job completed: %v\n", result)
// 	}

// 	fmt.Printf("it took %s ms.\n", time.Since(startTime))
// }

package main

import (
	"fmt"
	"sync"
	"time"
)

type Result struct {
	Value string
	Err   error
}

func worker(jobs chan string, wg *sync.WaitGroup, resultChan chan Result) {
	defer wg.Done()

	for job := range jobs {
		time.Sleep(time.Microsecond * 50)
		fmt.Println("image processed:\n", job)

		resultChan <- Result{
			Value: job,
			Err:   nil,
		}

	}
	fmt.Printf("Worker shutting Down!!\n")
}

func main() {
	var wg sync.WaitGroup

	jobs := []string{
		"img1.png",
		"img2.png",
		"img3.png",
		"img4.png",
		"img5.png",
		"img6.png",
		"img7.png",
		"img8.png",
		"img9.png",
		"img10.png",
	}
	totalWorkers := 10
	jobsChan := make(chan string, 50)
	resultChan := make(chan Result, len(jobs))

	for i := 1; i <= totalWorkers; i++ {
		wg.Add(1)
		go worker(jobsChan, &wg, resultChan)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for i := 0; i < len(jobs); i++ {
		jobsChan <- jobs[i]
	}

	close(jobsChan)

	for result := range resultChan {
		fmt.Printf("Received processed image:%v\n", result)
	}
}
