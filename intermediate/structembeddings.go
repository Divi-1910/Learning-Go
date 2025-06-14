package main

import "fmt"

type person struct {
	name string
	age  int
}

type Employee struct {
	person // embedded struct
	empId  string
	salary int
}

func (p person) introduce() {
	fmt.Printf("Hiii , I am %s , I am %d years old\n", p.name, p.age)
}

func (e Employee) introduce() {
	fmt.Println("Hii , I am ", e.name, "My Id is ", e.empId)
}

func main() {
	emp := Employee{
		person: person{"John", 25},
		empId:  "123",
		salary: 10000,
	}

	fmt.Println(emp)
	fmt.Println(emp.name, emp.age) // direct access to name
	emp.introduce()

}
