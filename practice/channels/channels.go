package main

// [1.]****** Receiving Data to Goroutine *****
// func processNum(numChan chan int) {
// 	for num := range numChan {
// 		fmt.Println("processing number", num)
// 		time.Sleep(time.Second)
// 	}
// }

// ***** [2.] From Goroutine, Sending Received Channel Data to Channel again
// func sum(result chan int, num1 int, num2 int) {
// 	sum := num1 + num2
// 	result <- sum
// }

// [3.]**** Goroutine Synchronization or Un-Buffered Channel *****
// func task(done chan bool) {
// 	defer func() {
// 		done <- true
// 	}()
// 	fmt.Println("Processin...")
// }

// [4.] ***** Buffered Channel: We can send limited amount of data without blocking. *****
// For Example we are implementing: Queue System for Email Sending:---
// func sendEmail(emailChan chan string, done chan bool) {
// 	defer func() {
// 		done <- true
// 	}()
// 	for email := range emailChan { // Here it's doing infinite loop because it's a buffered channel, so this is the blocked operation thats why we need to close the emailChan channel.
// 		fmt.Println("Sending email to:", email)
// 		time.Sleep(time.Second)
// 	}
// }

// [5.] ***** Buffered Channel and Listining on Multiple channel ******
// func listenChannels(ch1 chan int, ch2 chan string) {
// 	for i := 1; i <= 2; i++ {
// 		select {
// 		case v1 := <-ch1:
// 			fmt.Println("Received from ch1:", v1)
// 		case v2 := <-ch2:
// 			fmt.Println("Received from ch2:", v2)
// 		}
// 	}

// }

// func main() {
// [5.] ***** Buffered Channel and Listining on Multiple channel ******
// chan1 := make(chan int)
// chan2 := make(chan string)

// go func() {
// 	chan1 <- 10
// }()
// go func() {
// 	chan2 <- "ping"
// }()

// for i := 1; i <= 2; i++ {
// 	select {
// 	case chan1Val := <-chan1:
// 		fmt.Println("Received data from:", chan1Val)
// 	case chan2Val := <-chan2:
// 		fmt.Println("Received data from:", chan2Val)
// 	}
// }

// chan1 := make(chan int)
// chan2 := make(chan string)
// Start listeners
// go listenChannels(chan1, chan2)

// Send Messages
// chan1 <- 10
// chan2 <- "ping"

//[4.] **** Buffered Channel ******
// emailChan := make(chan string, 5)
// done := make(chan bool)

// go sendEmail(emailChan, done)

// for i := 1; i <= 5; i++ {
// 	emailChan <- fmt.Sprintf("%d@exmple.com", i)
// }
// fmt.Println("Done Sending.✅")
// This is important to close channels
// close(emailChan)
// fmt.Println(<-done)
// close(done)

// [3.]**** Goroutine Synchronization or Un-Buffered Channel *****
// ** Un-Buffered means: We can send only one data until it receive, we can't send another data

// done := make(chan bool)

// go task(done)
// We need to bolck here by: <- done (this <- done makes our program stop until when someone doesn't send data on done channel)
// <-done // block

// [2.]****** Getting channel data from Goroutine ******
// result := make(chan int)

// go sum(result, 2, 5)
// fmt.Println("Total sum:", <-result)
// [1.]****** Sending data to Goroutine ******
// numChan := make(chan int)
// go processNum(numChan)
// for {
// 	numChan <- rand.Intn(100)
// }

// messageChan := make(chan string)

// messageChan <- "Ping" // blocking

// data := <-messageChan
// fmt.Println(data)
// }
