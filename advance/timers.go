package main

import (
	"fmt"
	"time"
)

/*
Why use Timers?
	Timeouts
	Delays
	Periodic Tasks

The `time.Timer` Type
	Creating a Timer
	Timer Channel
	Stopping a timer

Practical Use cases for timers :
 	implementing timeouts
	scheduling delayed operations
	periodic tasks
	handle large numbers of goroutines
	use `defer` for unlocking

Best Practices
	Avoid Resource Leaks
	Combine with channels
	Use `time.After` for simple timeouts

*/

func main() {
	fmt.Println("Starting the app...")
	timer := time.NewTimer(2 * time.Second)
	fmt.Println("Waiting for the timer...")
	<-timer.C // blocking in nature
	fmt.Println("Timer expired")

	fmt.Println("Starting the app...")
	timer = time.NewTimer(2 * time.Second)
	fmt.Println("Waiting for the timer...")
	stopped := timer.Stop()
	if stopped {
		fmt.Println("Timer stopped")
	}
	// <-timer.C // blocking in nature
	// fmt.Println("Timer expired")

	fmt.Println("Starting the app...")
	timer = time.NewTimer(2 * time.Second)
	fmt.Println("Waiting for the timer...")
	stopped = timer.Stop()
	if stopped {
		fmt.Println("Timer stopped")
	}
	timer.Reset(time.Second)
	<-timer.C // blocking in nature
	fmt.Println("Timer expired")

	// Implementing timeouts
	timeOut := time.After(3 * time.Second)
	done := make(chan bool)

	go func() {
		longRunningOperation()
		done <- true
	}()

	select {
	case <-timeOut:
		fmt.Println("Operation Timed out")
	case <-done:
		fmt.Println("Operation done")
	}

	// Scheduling Delayed operations
	timer = time.NewTimer(2 * time.Second)
	go func() {
		<-timer.C
		fmt.Println("Delayed operation executed")
	}()

	fmt.Println("Waiting...")
	time.Sleep(3 * time.Second)

	timer1 := time.NewTimer(1 * time.Second)
	timer2 := time.NewTimer(2 * time.Second)
	timer1Expired := false
	timer2Expired := false

	for {
		select {
		case <-timer1.C:
			fmt.Println("Timer1 expired")
			timer1Expired = true
		case <-timer2.C:
			fmt.Println("Timer2 expired")
			timer2Expired = true
		}
		if timer1Expired && timer2Expired {
			break
		}
	}

	fmt.Println("End of the app...")

}

func longRunningOperation() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}
