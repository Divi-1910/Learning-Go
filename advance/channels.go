package main

import (
	"fmt"
	"time"
)

/*
	Why use Channels ??
		Enable safe and efficient communication between concurrent Goroutines
		Help Synchronize and manage the flow of data in concurrent programs

	Basics of Channels
		Creating Channels (make(chan type))
		Sending and Receiving Data (<-)
		Channel Directions
			Send Only : `ch <- value'
			Receive Only : `value := <-ch`

	Common Pitfalls and Best Practices
		Avoid Deadlocks
		Avoiding Unnecessary Buffering
		Channel Direction
		Graceful Shutdown
		Use defer for unlocking
*/

// go routines are non-blocking
// channels are blocking

func main() {
	//variable :=  make(chan type)

	greeting := make(chan string)
	greet := "Hello Welcome"

	go func() {
		time.Sleep(1 * time.Second)
		greeting <- greet
		time.Sleep(1 * time.Second)
		greeting <- "Divyansh"
	}()

	receiver1 := <-greeting // blocking
	receiver2 := <-greeting // blocking

	fmt.Println(receiver1)
	fmt.Println(receiver2)

	go func() {
		greeting <- greet
		greeting <- "Hiiii"
		for _, e := range "How are you?" {
			greeting <- "Alphabet : " + string(e)
		}
	}()

	receiver3 := <-greeting
	fmt.Println(receiver3)
	receiver4 := <-greeting
	fmt.Println(receiver4)

	for _, _ = range "how are you?" {
		receiver5 := <-greeting
		fmt.Println(receiver5)
	}

}
