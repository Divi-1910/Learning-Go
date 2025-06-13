package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (q int, r int) {
	q = a / b
	r = a % b
	return
}

func compare(a, b int) (string, error) {
	if a > b {
		return "a is greater than b", nil
	} else if a < b {
		return "a is less than b", nil
	} else {
		return "", errors.New("unable to Compare , as they are equal")
	}
}

func main() {
	// func functionName(paramter1 type1 , parameter2 type2,...) (returnType1 , returnType2,....){
	//	code block
	// return returnvalue1 , returnvalue2
	//}

	q, r := divide(10, 5)
	fmt.Println(q, r)

	result, error := compare(5, 5)
	if error != nil {
		fmt.Println("Error is :", error)
	}
	fmt.Println(result)

}
