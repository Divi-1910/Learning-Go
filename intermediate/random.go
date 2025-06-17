package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(rand.Intn(100) + 1)
	found := false
	for found != true {
		num := rand.Intn(100)
		fmt.Println(num)
		if num+1 == 1 {
			found = true
		}
	}

	val := rand.New(rand.NewSource(42))
	fmt.Println("Random Number is : ", val.Intn(6)+5) // this is will generate the same random number

	fmt.Println(rand.Float64()) // between 0.0 and 1.0

	for {
		fmt.Println("Welcome to the Dice Game !! ")
		fmt.Println("1. Roll the dice")
		fmt.Println("2. Exit")

		fmt.Print("Enter your Choice : ")

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil || (choice != 1 && choice != 2) {
			fmt.Println("Error reading input: ", err)
			continue
		}

		if choice == 2 {
			fmt.Println("Thanks for Playing , Bye Bye !! ")
			break
		}

		die1 := rand.Intn(6) + 1
		die2 := rand.Intn(6) + 1

		fmt.Println("You rolled a ", die1, "and a ", die2)
		fmt.Println("Total is : ", die1+die2)

		fmt.Print("Do you want to roll Again (y/n) : ")
		var rollAgain string
		_, err = fmt.Scan(&rollAgain)
		if err != nil || (rollAgain != "y" && rollAgain != "n") {
			fmt.Println("Error reading input: ", err)
			fmt.Println("Assuming No")
			break
		}
		if rollAgain == "n" {
			fmt.Println("Thanks for Playing , Bye Bye !! ")
			break
		} else if rollAgain == "y" {
			fmt.Println("Let's Roll Again !! ")
			continue
		}

	}

}
