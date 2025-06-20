package main

import (
	"flag"
	"fmt"
	"os"
)

/*
	Key Concepts
		1. `os.Args` Slice
		2. Parsing Arguments
		3. `flag` Package
		4. Default Values and Usage

	Considerations
		1. Order of Arguments
		2. Flag Reuse
		3. Order of Flags
		4. Default Values
		5. Help Output

	Best Practices
		1. Clear Documentation
		2. Consistent Naming
		3. Validation

*/

func main() {
	fmt.Println("Command: ", os.Args[0])

	// go run CLIArgs.go hello world go is cool -> command to run
	//for _, arg := range os.Args {
	//	fmt.Println("Argument : ", arg)
	//}

	// define flags

	var name string
	var age int
	var isMarried bool

	flag.StringVar(&name, "name", "Divyansh", "Name of the User")
	flag.IntVar(&age, "age", 18, "Age of the User")
	flag.BoolVar(&isMarried, "isMarried", false, "Is the person Married?")

	flag.Parse()

	fmt.Println("Name : ", name)
	fmt.Println("Age : ", age)
	fmt.Println("isMarried : ", isMarried)

}
