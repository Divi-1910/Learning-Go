package main

import "fmt"

type Person struct {
	Name    string
	Age     int
	address Address
	PhoneHomeCell
}

// method on a struct
func (p Person) fullname() string {
	return p.Name
}

func (p *Person) incrementAge() {
	p.Age += 1
}

type Address struct {
	city  string
	state string
}

type PhoneHomeCell struct {
	home string
	cell string
}

func main() {

	divyansh := Person{Name: "divyansh", Age: 22, address: Address{
		city:  "London",
		state: "United Kingdom",
	}, PhoneHomeCell: PhoneHomeCell{
		home: "123456789",
		cell: "0987654321",
	}}

	Sharvari := Person{Name: "Sharvari", Age: 22, address: Address{
		city:  "Pune",
		state: "United Kingdom",
	}}

	Sharvari.address.state = "India"

	fmt.Println(divyansh.Name)
	fmt.Println(divyansh.Age)
	fmt.Println(Sharvari)
	fmt.Println(divyansh.cell)
	fmt.Println(divyansh.home)

	fmt.Println("Can divyansh and sharvari similar? : ", divyansh == Sharvari)

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
