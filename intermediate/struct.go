package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// method on a struct
func (p Person) fullname() string {
	return p.Name
}

func (p *Person) incrementAge() {
	p.Age += 1
}

func main() {

	divyansh := Person{Name: "divyansh", Age: 22}
	fmt.Println(divyansh.Name)
	fmt.Println(divyansh.Age)

	// Annoynous Structs
	User := struct {
		username string
		email    string
	}{
		username: "john",
		email:    "johndoe@gmail.com",
	}
	fmt.Println(User)
	fmt.Println(divyansh.fullname())
	divyansh.incrementAge()
	fmt.Println(divyansh.Age)

}
