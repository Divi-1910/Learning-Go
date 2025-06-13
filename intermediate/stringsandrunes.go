package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	message := "Hello from Go!"
	rawMessage := fmt.Sprintf("Hello %s", message)

	fmt.Println(message)
	fmt.Println(rawMessage)

	fmt.Println(len("hello\ndone"))

	char := message[2]
	fmt.Println(string(char))

	for i, ch := range message {
		fmt.Printf("%d: %c\n", i, ch)
	}

	runestr1 := []rune("Hello , Greetings from me")
	runestr2 := 'H'

	newmessage := "Hello world from divyansh"

	for index, ch := range newmessage {
		fmt.Printf("%d: %c\n", index, ch)
	}

	fmt.Println(utf8.RuneCountInString(string(runestr1)))
	fmt.Println(utf8.ValidString(string(runestr2)))

}
