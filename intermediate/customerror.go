package main

import (
	"errors"
	"fmt"
)

type customError struct {
	code    int
	message string
}

// method on custom error
func (e *customError) Error() string {
	return fmt.Sprintf("Error %d : %s", e.code, e.message)
}

func doSomething() error {
	err := doSomethingElse()
	if err != nil {
		return &customError{
			code:    500,
			message: "Something Went Wrong",
			er:      err,
		}
	}
	return nil
}

func doSomethingElse() error {
	return &customError{
		code:    500,
		message: "Something Failed",
	}
}

func main() {
	err := doSomething()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("OK")
}
