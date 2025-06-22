package main

import (
	"fmt"
	"time"
)

/*
	Why are Channel Directions Important?
		Improve code clarity and maintainability
		Prevent Unintended Operations on channels
		Enhance type safety by clearly defining the channel's purpose

	Basic Concepts of Channel Directions
		Unidirectional Channels
		Send Only Channels
		Receive Only Channels
		Testing and Debugging

	Defining Channel Directions in Function Signatures
		Send Only parameters ( func produceData(ch chan<-int))
		Receive Only parameters (func consumeData(ch<-chan int))
		Bidirectional Channels (func bidirectional(ch chan int))

*/

func main() {
	ch := make(chan int)

	go func(ch chan<- int) {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}(ch)

	for value := range ch {
		fmt.Println("Received : ", value)
	}

	ch1 := make(chan int, 2)
	producer(ch1)
	consumer(ch1)

}

func producer(ch chan<- int) { // send only channel
	for i := 0; i < 5; i++ {
		fmt.Println("Producing : ", i)
		ch <- i
		time.Sleep(500 * time.Millisecond)
	}
	close(ch)
}

func consumer(ch <-chan int) { // receive only channel
	for value := range ch {
		fmt.Println("Consumed : ", value)
		time.Sleep(1000 * time.Millisecond)
	}
}
