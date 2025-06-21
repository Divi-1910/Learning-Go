package main

import (
	"fmt"
	"strconv"
	"time"
)

/*
Why is Channel Synchronization important?
	Ensures that data is properly exchanged between Goroutines
	Coordinates the execution flow to avoid race conditions and ensure predictable behaviour
	Helps manage the lifecycle of Goroutines and the completion tasks

Common Pitfalls and Best practices
	Avoid Deadlocks
	Close Channels
	Avoid Unnecessary Blocking

*/

func main() {
	done := make(chan struct{})

	go func() {
		fmt.Println("Working...")
		time.Sleep(3 * time.Second)
		done <- struct{}{}
	}()

	<-done
	fmt.Println("End of Program")

	chint := make(chan int)
	go func() {
		fmt.Println("Sending...")
		chint <- 1
		time.Sleep(1 * time.Second)
		fmt.Println("Value Sent") // this will not get executed
	}()
	value := <-chint
	fmt.Println(value)

	numGoroutines := 3
	doneint := make(chan int, numGoroutines)

	for i := range numGoroutines {
		time.Sleep(1 * time.Second)
		go func(id int) {
			fmt.Printf("Goroutine %d working...\n", id)
			time.Sleep(1 * time.Second)
			doneint <- id // sending signal of completion
		}(i)
	}

	for i := 0; i < numGoroutines; i++ {
		<-doneint // blocking until all the go routines are completed
	}

	// SYNCHRONIZATION DATA EXCHANGE
	data := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			data <- "Hello " + strconv.Itoa(i)
			time.Sleep(100 * time.Millisecond)
		}
		close(data)
	}()
	for value := range data {
		fmt.Println(value, time.Now())
	}

}
