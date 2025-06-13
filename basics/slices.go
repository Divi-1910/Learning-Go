package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	arr := []int{1, 2, 3}
	fmt.Println(arr)
	slice := make([]int, 5)
	arr1 := [5]int{1, 2, 3, 4, 5}
	slice1 := arr1[1:4]
	fmt.Println(slice1)
	slice = append(slice, 6, 7)
	fmt.Println(slice)
	sliceCopy := make([]int, len(slice))
	copy(sliceCopy, slice)
	fmt.Println("sliceCopy ", sliceCopy)

	//	fmt.Println(slice)
	for _, value := range slice1 {
		fmt.Println(value)
	}

	fmt.Println("Element at index 3 of slice1 , ", slice[2])

	if slices.Equal(slice1, slice) {
		fmt.Println("slice1 and slice2 are equal")
	} else {
		fmt.Println("slice1 and slice2 are not equal")
	}

	twoDim := make([][]int, 3)

	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoDim[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoDim[i][j] = i + j
		}
	}

	fmt.Println(twoDim)
	fmt.Println(cap(twoDim))
	fmt.Println(len(twoDim))

	temp := []int{5, 3, 7, 1}

	fmt.Println(temp)

}
