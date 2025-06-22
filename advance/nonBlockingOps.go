package main

import (
	"fmt"
	"time"
)

/*
Why use Non-Blocking Operations?
	Avoid Deadlocks
	Improve Efficiency
	Enchance Concurrency

Best Practices for Non-Blocking Operations
	Avoid Busy Waiting
	Handle Channel Closure Properly
	Combine with Contexts for Cancellations
	Ensure Channel Capacity Management

*/

func main() {
	ch := make(chan int)

	// non-blocking receive operation
	select {
	case msg := <-ch:
		fmt.Println("Received : ", msg)
	default:
		fmt.Println("No message received")
	}

	// non-blocking send operation
	select {
	case ch <- 1:
		fmt.Println("Message Sent")
	default:
		fmt.Println("Channel is not ready to receive")
	}

	// non-blocking operation in real time systems
	data := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case d := <-data:
				fmt.Println("Received : ", d)
			case <-quit:
				fmt.Println("Stopping...")
				return
			default:
				fmt.Println("Waiting for data...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	for i := 0; i < 5; i++ {
		data <- i
		time.Sleep(1000 * time.Millisecond)
	}
	quit <- true

}
