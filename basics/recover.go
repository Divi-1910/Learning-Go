package main

import "fmt"

func process() {
	defer func() {
		// if r:=recover(): r!=nil {

		r := recover()
		if r != nil {
			fmt.Println("Recovered from panic :", r)
		}
	}()
	fmt.Println("start process")
	panic("Something bad happened")
	fmt.Println("end process")
}

func main() {
	// recover is a built-in funcion in go used to regain control
	// of a panicking goroutine
	// It's used inside a defer function to stop panic and prevent
	// the program from crashing

	process()
	fmt.Println("Return From Process")

	/*
		Practical Use Cases of recover
			1. Graceful Recovery
			2. Cleanup
			3. Logging and Reporting

		Best Practices
			1. Use with defer
			2. Avoid Silent Recovery
			3. Avoid Overuse

	*/

}
