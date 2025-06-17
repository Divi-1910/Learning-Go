package main

import "fmt"

func main() {
	num := 123456789
	fmt.Printf("%05d\n", num)

	message := "Hello"
	fmt.Printf("|%-10s|\n", message)

	message1 := "Hello \n World"
	message2 := `Hello \n World` // back ticks

	fmt.Println(message1)
	fmt.Println(message2)

}
