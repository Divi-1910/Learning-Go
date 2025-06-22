package main

import (
	"fmt"
	"time"
)

/*
Why close channels?
	signal completion
	prevent resource leaks

Best Practices for closing channels
	close channels only from the sender
	avoid closing channels more than once
	avoid closing channels from multiple goroutines


*/

func main() {

	// simple closing channel example
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	for value := range ch {
		fmt.Printf("Received from Channel %d\n", value)
	}

	fmt.Println()

	ch1 := make(chan int)
	close(ch1)

	val, ok := <-ch1
	if !ok {
		fmt.Println("Channel 1 closed")
	} else {
		fmt.Println("Received from Channel 1 : ", val)
	}

	fmt.Println()

	ch2 := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch2 <- i
		}
		close(ch2)
	}()

	for value := range ch2 {
		fmt.Printf("Received from Channel 2 : %d\n", value)
	}

	fmt.Println()
	// CLosing channels can only be done by the sender
	// not good practices in closing channels
	ch3 := make(chan int)
	go func() {
		close(ch3)
		// close(ch3) -> panic occurs when closing a closed channel
	}()
	time.Sleep(1000 * time.Millisecond)

	fmt.Println()

	// PIPELINE PATTERN FOR CHANNELS
	ch4 := make(chan int)
	ch5 := make(chan int)

	go producer(ch4)
	go filter(ch4, ch5)

	for value := range ch5 {
		fmt.Printf("Received from Channel 5 : %d\n", value)
	}

}

func producer(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}

func filter(in <-chan int, out chan<- int) {
	for value := range in {
		if value%2 == 0 {
			out <- value
		}
	}
	close(out)
}
