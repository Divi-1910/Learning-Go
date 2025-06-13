package main

import "fmt"

func main() {

	// a pointer is a variable that stores the memory
	// address of another variable

	// Use Cases
	// Modify the value of a variable indirectly
	// pass large data structures efficiently between functions
	// manage memory directly for performance reasons

	// var ptr * int (ptr is a pointer to an integer)

	var ptr *int
	a := 10
	ptr = &a
	fmt.Println(ptr)
	fmt.Println(a)
	*ptr++
	fmt.Println(a)
	fmt.Println(*ptr)

	// pointer arthimetic is not supported in go

}
