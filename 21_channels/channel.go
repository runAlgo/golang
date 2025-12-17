package main

import (
	"fmt"
	"time"
)

// ****[4] Buffered Channels ****
// If you want to implementing something interesting like QUEUE System,
// then use Buffered Channel.
// In Buffered channel we can able to send a limited block of data without blocking.

func emailSender(emailChan <-chan string, done chan<- bool) { 
	// <-chan  **This means we only receive data on channel
	// chan<-  **This means we only send data on channel
	
	for email := range emailChan {
		defer func() { done <- true }()
		fmt.Println("Sending email to:", email)
		time.Sleep(time.Second)
	}
}

//[2] ***** Sender channel *****
// func processNum(numChan chan int) {
// 	for num := range numChan {
// 		fmt.Println("processing number", num)
// 		time.Sleep(time.Second)
// 	}
// }

// [3]**** Receiver Channel *****
// func sum(result chan int, num1, num2 int) {
// 	sum := num1 + num2
// 	result <- sum
// }

//[1] ***** goroutine synchronizer ****
// func task(done chan bool) {
// 	defer func () { done <- true }()
// 	fmt.Println("Processing...")
// }

func main() {
	// [5]***** Listening to multiple channels at a time *****
	// chan1 := make(chan int)
	// chan2 := make(chan string)

	// go func() {
	// 	chan1 <- 10
	// }()

	// go func() {
	// 	chan2 <- "pong"
	// }()

	// for i := 0; i < 2; i++ {
	// 	select {
	// 	case chan1Val := <-chan1:
	// 		fmt.Println("Received data from chan1:", chan1Val)
	// 	case chan2Val := <-chan2:
	// 		fmt.Println("Received data from chan2:", chan2Val)
	// 	}

	// }

	// [4]****** Buffered Channel ********
	emailChan := make(chan string, 5)
	done := make(chan bool)

	go emailSender(emailChan, done)

	for i := 1; i <= 5; i++ {
		emailChan <- fmt.Sprintf("%dexample@gmail.com", i)
	}

	fmt.Println("Done sending.")
	// *** this is important we need to close the buffered channel *****
	close(emailChan)
	<-done

	// emailChan <- "01@gmail.com"
	// emailChan <- "02@gmail.com"

	// fmt.Println(<-emailChan)
	// fmt.Println(<-emailChan)

	//[1]**** Un-buffered Channel ****
	// If you want to do task one by one then you this normal un-buffered channel.
	// If you want to send data as-well-as if you want to wait for processing that sent data, then use un-buffered or Normal channel.
	// done := make(chan bool)
	// We can send only one data at a time, We can't able send another data if previous sent data doesn't received to another side.

	// go task(done)
	// <-done // block

	//[2]**** Sender Channel ****
	// numChan := make(chan int)

	// go processNum(numChan)

	// for {
	// 	numChan <- rand.Intn(100)
	// }

	//[3]**** Receiver Channel ****
	// result := make(chan int)

	// go sum(result, 4, 5)

	// res := <-result
	// fmt.Println("Sum is:", res)

	// time.Sleep(time.Second * 2)

	// messageChan := make(chan string)

	// messageChan <- "ping" // send data to messageChan

	// msg := <-messageChan // receive data from messageChan

	// fmt.Println(msg)
}
