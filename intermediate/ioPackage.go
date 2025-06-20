package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

/*

Why is the io package important?
	1. Facilitates interaction with various data sources (files , networks , in-memory buffers)
	2. Provides a consistent interface for handling I/O Operations

Core Interfaces :
	1. io.Reader
	2. io.Writer
	3. io.Closer

Common Types and Functions
	1. io.Reader
 	2. io.Writer
	3. io.Copy()
	4. io.MultiReader()
	5. io.Pipe()

Working with Buffers
	1. bytes.Buffer
	2. bufio Package

*/

func readFromReader(r io.Reader) {
	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	if err != nil {
		log.Fatalln("Error reading from reader:", err)
		return
	}
	fmt.Println("Read Data : ", string(buf[:n]))

}

func writeToWriter(writer io.Writer, data string) {
	_, err := writer.Write([]byte(data))
	if err != nil {
		log.Fatalln("Error writing to writer:", err)
		return
	}
}

func closeResource(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatalln("Error closing resource:", err)
		return
	}
}

func bufferExample() {
	var buf bytes.Buffer // is bytes.buffer (value type) (in stack memory)
	buf.WriteString("Hello Buffer !!")
	fmt.Println(buf.String())
}

func multiReaderExample() {
	r1 := strings.NewReader("Hello")
	r2 := strings.NewReader("World")
	mr := io.MultiReader(r1, r2)
	buf := new(bytes.Buffer) // pointer to bytes.Buffer (pointer type) (in heap memory)
	_, err := buf.ReadFrom(mr)
	if err != nil {
		log.Fatalln("Error reading from reader:", err)
		return
	}
	fmt.Println("Read Data : ", buf.String())
}

func PipeExample() {
	pr, pw := io.Pipe()
	go func() {
		pw.Write([]byte("Hello Pipes !! "))
		pw.Close()
	}()

	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(pr)
	if err != nil {
		return
	}
	fmt.Println("Read Data : ", buf.String())

}

func writeToFile(filename string, data string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("Error opening file:", err)
		return
	}
	defer closeResource(file)

	_, err = file.Write([]byte(data))
	if err != nil {
		log.Fatalln("Error writing to file:", err)
	}

	writer := io.Writer(file)
	_, err = writer.Write([]byte(data))
	if err != nil {
		log.Fatalln("Error writing to file:", err)
	}

}

func main() {
	fmt.Println("=== Read from Reader ===")
	readFromReader(strings.NewReader("Hello From Divyansh"))

	fmt.Println("=== Write to Writer ===")
	var writer bytes.Buffer
	writeToWriter(&writer, "Hello World")
	fmt.Println(writer.String())

	fmt.Println("=== Buffer Example ===")
	bufferExample()

	fmt.Println("=== MultiReader Example ===")
	multiReaderExample()

	fmt.Println("=== PipeExample ===")
	PipeExample()

	filepath := "io.txt"
	writeToFile(filepath, "Hello File")

}
