package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	str := "Hello God"
	fmt.Println(str)

	sz := len(str)
	fmt.Println(sz)

	str1 := "Hello"
	str2 := "World"

	str3 := str1 + " " + str2
	fmt.Println(str3)

	fmt.Println(str1[0])         // this returns ascii value
	fmt.Println(string(str1[0])) // this returns the character

	fmt.Println(str3[1:7])

	// integer to string
	num := 23
	str4 := strconv.Itoa(num)
	fmt.Println(str4)
	fmt.Println(len(str4))

	fruits := "Apple,orange,banana"
	parts := strings.Split(fruits, ",")
	fmt.Println(parts)

	countries := []string{"America", "Europe", "India", "Japan", "USA"}
	joinedCountires := strings.Join(countries, ",")
	fmt.Println(joinedCountires)

	fmt.Println(strings.Contains(joinedCountires, "India"))

	replaced := strings.Replace(joinedCountires, "India", "Bharat", 1)
	fmt.Println(replaced)

	str_with_space := "            Hello Strings!!!!!!!!!               "
	fmt.Println(str_with_space)
	fmt.Println(strings.TrimSpace(str_with_space))

	fmt.Println(strings.ToLower(fruits))
	fmt.Println(strings.ToUpper(fruits))

	fmt.Println(strings.Repeat(fruits, 5))

	fmt.Println(strings.Count("Hello", "1"))
	fmt.Println(strings.HasPrefix("Hello", "H"))
	fmt.Println(strings.HasSuffix("Hello", "e"))

	str5 := "Hello, 123 Go 456"
	re := regexp.MustCompile(`\d+`)
	fmt.Println(re)
	fmt.Println(str5)
	matches := re.FindAllString(str5, -1)
	fmt.Println(matches)

	str6 := "Hello , HHH"
	fmt.Println(utf8.RuneCountInString(str6))

}
