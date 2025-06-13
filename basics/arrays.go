package main

import (
	"fmt"
	"sort"
)

func main() {
	// var arrayName [size]elementType
	var nums [5]int
	fmt.Println(nums)

	nums[0] = 1
	nums[4] = 4

	fmt.Println(nums)

	fruits := [3]string{"apple", "banana", "orange"}
	fmt.Println(fruits)

	originalArray := [5]int{1, 2, 3, 4, 5}

	copiedArray := originalArray

	originalArray[0] = 10

	fmt.Println(originalArray)

	fmt.Println(copiedArray)

	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i], " -- ", i)
	}

	for index, value := range fruits {
		fmt.Println(index, " -- ", value)
	}

	for _, num := range originalArray {
		fmt.Println(num)
	}

	_, b := solve()
	fmt.Println("b :", b)

	array1 := [3]int{1, 2, 3}
	array2 := [3]int{1, 2, 3}

	fmt.Println(array1 == array2)

	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(matrix)
	var copiedArray3 *[3][3]int
	copiedArray3 = &matrix
	copiedArray3[0][1] = 1000
	fmt.Println(matrix)

}

func solve() (int, int) {
	return 1, 2
}
