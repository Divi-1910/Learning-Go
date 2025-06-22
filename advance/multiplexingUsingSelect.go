package main

import (
	"fmt"
	"time"
)

/*
Why use multiplexing??
	Concurrency
	Non-blocking Operations
	Timeouts and Cancellations

Best Practices for using `select`
	Avoiding Busy Waiting
	Handling Deadlocks
	Readability and Maintainability
	Testing and Debugging

*/

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- 10
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- 20
	}()

	time.Sleep(2 * time.Second)
	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			fmt.Println("Received from Channel 1 : ", msg)

		case msg := <-ch2:
			fmt.Println("Received from Channel 2 : ", msg)

		default:
			fmt.Println("No message received")
		}
	}

	ch3 := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		ch3 <- 30
		close(ch3)
	}()

	select {
	case msg := <-ch3:
		fmt.Println("Received from Channel 3 : ", msg)
	case <-time.After(3 * time.Second):
		fmt.Println("Timed out")
	}

	ch4 := make(chan int)
	go func() {
		ch4 <- 40
		close(ch4)
	}()

	for {
		select {
		case msg, ok := <-ch4:
			if !ok {
				fmt.Println("Channel 4 closed")
				return
			}
			fmt.Println("Received from Channel 4 : ", msg)
		}
	}

	fmt.Println("End of Program")

}
