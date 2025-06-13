package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	Age       int
}

func main() {
	const MAX_RETRIES = 5

	var employeeID = 1001

	fmt.Println("Employee ID:", employeeID)

}
