package main

import (
	"fmt"
	"os"
)

/*
	Why use temporary files and directories?
		temporary storage
		isolation
		automatic cleanup
		default values and usage


	Best Practices :
		Security
		Naming

	Practical Applications :
		File Processing
		Testing
		Caching

	Considerations :
		Platform Differences
		Concurrency
*/

func main() {
	tempFile, err := os.CreateTemp("", "tempFile")
	if err != nil {
		fmt.Println("Error creating temp tempFile:", err)
		return
	}
	defer tempFile.Close()
	defer os.Remove(tempFile.Name())
	fmt.Println("Temporary tempFile name : ", tempFile.Name())

	tempDir, err := os.MkdirTemp("", "GoBootCamp")
	if err != nil {
		fmt.Println("Error creating temp tempDir:", err)
		return
	}
	fmt.Println("Temporary Directory Name : ", tempDir)

}
