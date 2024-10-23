package main

import (
	"fmt"
	// "time"
	// "sync"
)

// func printString(message chan string, wg *sync.WaitGroup) {
// 	fmt.Println(<-message)
// 	defer wg.Done()
// }

// func sumFun(result chan int, num1 int, num2 int){
// 	newRes := num1 + num2
// 	result <- newRes
// }

/*
func emailSender(emailChan chan string, doneChan chan bool) {
	defer func() { doneChan <- true }()
	for email := range emailChan {
		fmt.Println(email)
		time.Sleep(time.Second)
	}

}

func buffered_channel() {
	emailChan := make(chan string, 10)
	doneChan := make(chan bool)
	for i := 1; i <= 10; i++ {
		emailChan <- fmt.Sprintf("%d@email.com", i)
	}
	fmt.Println("Data is sent in the Channels")
	go emailSender(emailChan, doneChan)
	close(emailChan)
	<-doneChan
}

*/

func recieveFromMultipleChannels() {
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 1
	}()

	go func() {
		chan2 <- "string data is sent"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("Received from Channel 1", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("Recieved from Channel 2", chan2Val)
		}
	}

}

func main() {
	// var wg sync.WaitGroup
	// message := make(chan string)
	// wg.Add(1)
	// // sending data to go routine
	// go printString(message, &wg)
	// message <- "This is the data from main thread"

	// // receving data from channel
	// result := make(chan int)
	// go sumFun(result, 100,200)
	// res :=  <-result
	// fmt.Println("Sum :", res)
	// wg.Wait()

	// buffered Channel
	// buffered_channel()

	// receive data from multiple channels
	recieveFromMultipleChannels()
}
