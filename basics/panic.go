package main

import "fmt"

func process(input int) {
	if input < 0 {
		defer fmt.Println("Deferred Input Value :", input)
		panic("Invalid input")
	}
	fmt.Println("Processing : ", input)
}

func main() {
	// panic(interface {})

	// panic immediately stops normal execution flow of the program
	// and then all the deferred functions are still executed before the program crashes (emptying the stack trace)

	process(5)
	process(-1)

}
