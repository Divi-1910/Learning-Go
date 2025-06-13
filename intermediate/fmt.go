package main

import "fmt"

func main() {
	//fmt.Println("Hello world") // print with new line
	//fmt.Println('r')
	//fmt.Println(1234567.324324)
	//
	//fmt.Print("Divyansh ") // standard print
	//fmt.Print("Muley ")
	//
	//fmt.Printf("%s", "Hello from me ") // print with formatting verbs
	//
	//fmt.Println()
	//
	//// format to string
	//s := fmt.Sprint("Hello", "world", 123, 456, 789)
	//fmt.Println(s)
	//
	//// Scanning Functions
	//var name string
	//var age int
	//
	//fmt.Println("Enter your Name : ")
	//_, err := fmt.Scanln(&name)
	//if err != nil {
	//	return
	//}
	//fmt.Println("Enter your Age : ")
	//_, err = fmt.Scanln(&age)
	//if err != nil {
	//	return
	//}

	// Error Formatting Functions

	error1 := checkAge(15)
	if error1 != nil {
		fmt.Println("Error : ", error1)
	}

}

func checkAge(age int) error {
	if age < 18 {
		return fmt.Errorf("Age: %d is less than 18", age)
	}
	return nil
}
