package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("output.txt")
	if err != nil {
		fmt.Println("Error Opening file : ", err)
		return
	}
	defer func() {
		fmt.Println("Closing Open File")
		file.Close()
	}()
	fmt.Println("File Opened Successfully")

	data := make([]byte, 1024)
	_, err = file.Read(data)
	if err != nil {
		fmt.Println("Error Reading file : ", err)
		return
	}
	fmt.Println("Data : ", string(data))

	// the file pointer is at the end of the file now , we need to reset it
	_, err = file.Seek(0, 0)
	if err != nil {
		fmt.Println("Error Seeking file : ", err)
		return
	}

	scanner := bufio.NewScanner(file)
	// Read line by line
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Line : ", line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error Scanning file : ", err)
		return
	}

}
