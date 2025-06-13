package main

import (
	"fmt"
	"os"
)

func main() {

	defer fmt.Println("Deferred Statement")

	fmt.Println("Starting the main function ")
	os.Exit(1)
	fmt.Println("This will never be called")
}
