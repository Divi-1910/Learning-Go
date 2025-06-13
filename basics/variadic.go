package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, v := range nums {
		total = total + v
	}
	return total
}

func main() {

	//func functionName(param1 type1, param2 type2, param3.... type3) returnType{
	//	code block
	//}

	fmt.Println(sum(1, 2, 3, 4, 5))

}
