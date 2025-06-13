package main

import "fmt"

func main() {
	message := "hello world"
	runes := []rune(message)

	for i, v := range runes {
		if runes[i] == ' ' {
			continue
		}
		fmt.Printf("%d - %c \n", i, v)
		runes[i] = v + 1
	}

	fmt.Println(string(runes))
}
