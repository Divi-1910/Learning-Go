package main

import "fmt"

func main() {
	// defer in go is used to postpone the execution of a function until just before the surrounding function returns
	process(5)

	/*
		Practical Use cases of Defer
			1. Resource CleanUP
			2. Unlocking Mutexes
			3. Logging and Tracing

		Best Practices
			1. Keep Deferred Actions Short
			2. Understand Evaluation Timing
			3. Avoid Complex Control Flow
	
	*/

}

func process(i int) {
	defer fmt.Println("Deferred I Value :", i)
	defer fmt.Println("Deferred Statement Executed 1")
	defer fmt.Println("Deferred Statement Executed 2")
	defer fmt.Println("Deferred Statement Executed 3")
	i++
	fmt.Println("Normal Execution Statement Executed")
	fmt.Println("Actual Value of I is : ", i)
}
