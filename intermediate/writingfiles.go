package main

import (
	"fmt"
	"os"
)

func main() {
	// key components
	// 'os' package
	// Functions
	// Create(name string) (*File , error)
	// OpenFile(name string, flag int , perm FileMode) (*File , error)
	// Write(b []byte) (n int , err error)
	// WriteString(s string) (n int, err error)
	// best practices
	// error handling
	// deferred closing
	// permissions
	// buffering

	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error in Creating file : ", err)
		return
	}
	defer file.Close()

	// write data to file
	data := []byte("Writing in File \n")
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error in Writing to file : ", err)
		return
	}
	fmt.Println("Successfully Written data to the file output.txt")

	file, err = os.Create("writeString.txt")
	if err != nil {
		fmt.Println("Error in Creating file : ", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("Hello to file writing in go!! \n")
	if err != nil {
		fmt.Println("Error in Writing to file : ", err)
		return
	}
	fmt.Println("Successfully Written data to the file writeString.txt ")

}
