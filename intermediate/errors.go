package main

import (
	"errors"
	"fmt"
	"math"
)

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("math: square root of a negative number")
	}
	rt := math.Sqrt(x)
	return rt, nil
}

func process(data []byte) error {
	if len(data) == 0 {
		return errors.New("error: no data")
	}
	return nil
}

type myError struct {
	message string
}

func (e *myError) Error() string {
	return fmt.Sprintf("myError: %s", e.message)
}

func eprocess() error {
	return &myError{message: "myError"}
}

func main() {
	result1, err := sqrt(5)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result1)

	//result, err := sqrt(-5)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(result)

	// data := []byte{}
	if err := eprocess(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Data processed Successfully")

	err = readData()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Data read Successfully")
}

func readData() error {
	err := readConfig()
	if err != nil {
		return fmt.Errorf("readConfig: %w", err)
	}
	return nil
}

func readConfig() error {
	return errors.New("readConfig Error")
}
