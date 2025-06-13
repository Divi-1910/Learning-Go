package main

import "fmt"

func factorial(i int) int {
	if i == 0 {
		return 1
	}
	return i * factorial(i-1)
}

func SumOfDigits(n int) int {
	if n < 10 {
		return n
	}
	return n%10 + SumOfDigits(n/10)
}

func main() {
	fmt.Println(factorial(5))
	fmt.Println(SumOfDigits(1234))
}
