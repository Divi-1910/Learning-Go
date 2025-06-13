package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	target := random.Intn(100) + 1

	fmt.Println("Welcome to the guessing game")
	fmt.Println("Guess the number between 1 and 100")
	fmt.Println("Are you ready?")

	var guess int
	for {
		fmt.Println("Enter your guess : ")
		fmt.Scan(&guess)
		if guess == target {
			fmt.Println("Congratulations! You guessed the number")
			break
		} else if guess < target {
			fmt.Println("Too low")
		} else {
			fmt.Println("Too high")
		}

	}

}
