package main

import (
	"fmt"
	"time"
)

/*
	Why use Buffered Channels?
		1. Asynchronous Communication
		2. Load Balancing
		3. Flow Control

	Creating Buffered Channels
		1. make(chan Type, capacity)
		2. Buffer Capacity

	Key Concepts of Channel Buffering
		1. Blocking Behaviour
		2. Non-Blocking Operations
		3. Impact on Performance

	Best Practices for Using Buffered Channels
		1. Avoid Over-Buffering
		2. Graceful Shutdown
		3. Monitoring Buffer Usage
*/

func main() {
	// make(chan Type , capacity)
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Received : ", <-ch)
	}()
	ch <- 3
	fmt.Println("Value : ", <-ch)
	fmt.Println("Value : ", <-ch)

	ch1 := make(chan int, 2)
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- 1
	}()
	fmt.Println("Value from channel 1 : ", <-ch1)
	fmt.Println("End of Program")

	ch2 := make(chan int, 2)
	ch2 <- 1
	ch2 <- 2
	fmt.Println("Receiving from Buffer ")
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Received : ", <-ch2)
	}()
	ch2 <- 3
}
