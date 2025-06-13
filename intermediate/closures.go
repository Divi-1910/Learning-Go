package main

import "fmt"

func adder() func() int {
	i := 0
	fmt.Println("Previous Value of i : ", i)
	return func() int {
		i++
		fmt.Println("Adding 1 to i")
		return i
	}
}

func main() {
	// Closures and Scopes
	// a closure in go is defined inside another function
	// which captures and remembers variables from its
	// surrounding scope
	// even after the outer function has finished executing

	// Closures are a powerful feature for writing functions with a persistent
	// state , callbacks , and function factories

	//sequence := adder()
	//
	//fmt.Println(sequence())
	//fmt.Println(sequence())
	//fmt.Println(sequence())
	//fmt.Println(sequence())

	adderMain := func() func(int) int {
		sum := 0
		return func(x int) int {
			sum += x
			return sum
		}
	}()

	fmt.Println(adderMain(1))
	fmt.Println(adderMain(2))
	fmt.Println(adderMain(3))
	fmt.Println(adderMain(4))

}
