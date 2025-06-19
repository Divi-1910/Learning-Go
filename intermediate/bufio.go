package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	/* Bufio.Reader
	func NewReader(rd io.Reader) * Reader
	func (r *Reader) Read(p []byte) (n int , err error)
	func (r * Reader) ReadString(delim byte) (line string , err error)
	Bufio.Writer
		func NewWriter(wr io.Writer) *Writer
		func (w * Writer) write(p []byte) (n int , err error)
		func (w * Writer) writeString (s string) (n int , err error)
	Use Cases and Benefits
		1. Buffering
		2. Convenience Methods
		3. Error Handling
	Best Practices
		1. Check Errors
		2. Wrap Reader and Writer Instances for efficient I/O operations
		3. Don't Forget to call flush */

	reader := bufio.NewReader(strings.NewReader("Hello , Bufio Packageeee \n how you doing?"))

	// Reading
	data := make([]byte, 20)
	n, err := reader.Read(data)
	if err != nil {
		fmt.Println("Error Reading :", err)
		return
	}
	fmt.Printf("Read %d bytes %s \n", n, data[0:n])

	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error Reading :", err)
		return
	}
	fmt.Printf("Read %d bytes %s \n", n, line) // it started reading from the point it left

	writer := bufio.NewWriter(os.Stdout)
	data = []byte("Hello , bufio package!! \n")
	n, err = writer.Write(data)
	if err != nil {
		fmt.Println("Error Writing :", err)
		return
	}
	fmt.Printf("write %d bytes\n", n)
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error Flush :", err)
		return
	}

	str := "This is a string \n"
	n, err = writer.WriteString(str)
	if err != nil {
		fmt.Println("Error Writing :", err)
		return
	}
	fmt.Printf("wrote %d bytes\n", n)
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error Flush :", err)
		return
	}

}
