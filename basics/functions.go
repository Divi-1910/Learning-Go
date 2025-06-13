package main

import (
	"fmt"
)

func add(a float32, b int) float32 {
	return a + float32(b)
}

func applyOperation(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

func Multiplier(x int, y int) func(int, int) int {
	return func(x int, y int) int {
		return x * y
	}
}

func main() {
	// func <name> (parameters list) (returnTypes){
	// code block
	// return value/values
	// }

	// public functions start with a uppercase letter
	// private functions start with a lowerCase letter

	sum := add(10, 20)
	fmt.Println(sum)

	greet := func() {
		fmt.Println("Hello from anonymous function")
	}

	greet()

	operation := add
	ans := operation(2, 3)

	fmt.Println(ans)

	result := applyOperation(1, 2, multiply)

	fmt.Println(result)

	multiplied := Multiplier(5, 2)
	fmt.Println(multiplied(2, 3))

}

func multiply(x int, y int) int {
	return x * y
}
