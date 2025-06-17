package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("He Said ,\" I am here\" ")
	fmt.Println(`He said , "I am here"`)

	re := regexp.MustCompile(`[a-zA-Z0-9._+%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

	email1 := "user@email.com"
	email2 := "invalid_value"

	fmt.Println("Email1 : ", re.MatchString(email1))
	fmt.Println("Email2 : ", re.MatchString(email2))

	// Capturing Groups
	// compile a regex pattern to capture date components
	re = regexp.MustCompile(`((\d){4})-(\d{2})-(\d{2})`)

	date_str := "2024-07-30"

	submatches := re.FindStringSubmatch(date_str)
	fmt.Println(submatches)

	str := "Hello World"
	re = regexp.MustCompile(`[aeiou]`)
	result := re.ReplaceAllString(str, "*")
	fmt.Println(result)

	// i - case insensitive
	// m - multi line model
	// s - dot matches all

	re = regexp.MustCompile(`(?i)go`)
	test_str := "Golang is going great"
	fmt.Println("Match : ", re.MatchString(test_str))
}
