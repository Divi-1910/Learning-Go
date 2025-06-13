package main

import "fmt"

func main() {

	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for index, value := range numbers {
		fmt.Println("Index ", index, "Value ", value)
	}

	num := 1

	for num <= 5 {
		fmt.Print("Iteration ", num, "\n")
		num++
	}

}
