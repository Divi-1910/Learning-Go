package main

import (
	"fmt"
	"time"
)

/*
Why use Goroutines?
	simplify concurrent programming
	efficiently handle parallel tasks such as I/O operations, calculations , and more
	provide a way to perform tasks concurrently without manually managing threads

Basics of Goroutines
	Creating Goroutines (use the go keyword to start a new Goroutine)
	Goroutine LifeCycle
	Goroutine Scheduling

Goroutine Scheduling in Go
	Managed by the go runtime scheduler
	Uses M:N Scheduling Model
	Efficient Multiplexing

Common Pitfalls and Best Practices
	Avoiding Gorouting leaks
	Limiting Goroutine Creation
	Proper Error Handling
	Synchronization

*/

// Goroutines are just functions that leave the main thread and run in the background and come back to the join the main thread once the functions are finished/ready to return any value

// Goroutines do not stop the program flow and are non-blocking

func sayHello() {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello From Goroutine")
}

func main() {
	var err error

	fmt.Println("Beginning Program")
	go sayHello()
	fmt.Println("After sayHello function call")

	go func() {
		err = doWork()
	}()

	go printNumbers()
	go printLetters()

	time.Sleep(2 * time.Second)

	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		fmt.Println("Work Completed Successfully")
	}

}

func printNumbers() {
	for i := 0; i < 5; i++ {
		fmt.Println(time.Now(), "Number : ", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {
	for _, letter := range "abcdefghijklmnopqrstuvwxyz" {
		fmt.Println(time.Now(), "Letter : ", string(letter))
		time.Sleep(200 * time.Millisecond)
	}
}

func doWork() error {
	time.Sleep(1 * time.Second)
	fmt.Println("Working ...")
	return fmt.Errorf("An error Occurred while working")
}
