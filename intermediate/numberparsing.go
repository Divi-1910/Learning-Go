package main

import (
	"fmt"
	"strconv"
)

func main() {
	numStr := "12345"
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(num)
	fmt.Println(num + 1)

	num1, err := strconv.ParseInt(numStr, 10, 32)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(num1)

	floatStr := "3.1415926"
	floatval, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%.2f \n", floatval)

	binaryStr := "01010101010"
	binaryval, err := strconv.ParseInt(binaryStr, 2, 32)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binaryval)

}
