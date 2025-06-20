package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Key Components
	// Reading Lines Individually
	// Applying criteria or transformations
	// Processing each line
	// filtering lines based on content
	// removing empty lines
	// transforming line content
	// filtering lines by length
	// 'bufio' Package
	// Scanner.scan()
	// scanner.text()

	// Best Practices
	// Efficiency
	// Error Handling
	// Input Sources
	// Flexibility

	// Practical Applications
	// Data Applications
	//	Text Processing
	//	Data analysis

	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error in Opening file - example.txt : ", err)
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("Error in Closing file : ", err)
		}
	}()

	lineNumber := 1
	scanner := bufio.NewScanner(file)
	keyword := "important"

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, keyword) {
			updatedLine := strings.ReplaceAll(line, keyword, "valuable")
			fmt.Println(lineNumber, "Updated Line found : ", updatedLine)
			lineNumber++
			fmt.Println(lineNumber, "Previous Line Found:", line)
			lineNumber++
		}
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Error in Scanning file : ", err)
		return
	}

}
