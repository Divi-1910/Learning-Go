package main

import (
	"encoding/json"
	"fmt"
)

/*
	JSON - Javascript Object Notation
	`json.Marshal` - convert go data structures into Json (encoding)
	`json.Unmarshal` - convert json into go data structures (decoding)

	Best Practices
		Use Json tags
			Mapping Struct Fields to json keys
			Omitting Fields - if empty (`omitempty`) or always (`-`)
			Renaming Fields
			Controlling Json Encoding/Decoding Behaviour
		Validate Json
		use `omitempty`
		handle errors
		Custom Marshalling/Unmarshalling
*/

type Human struct {
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Age       int     `json:"age,omitempty"`
	Email     string  `json:"email"`
	Address   Address `json:"address"`
}

type Address struct {
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
}

func main() {
	human := Human{FirstName: "Divyansh", LastName: "Muley", Email: "dvmuley10@gmail.com", Address: Address{City: "Pune", State: "MH", Country: "India"}}

	jsonData, err := json.Marshal(human)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(string(jsonData))

	jsonData1 := `{
		"name" : "Sharvari",
		"EmployeeId" : 12345678,
		"address" : {
			"city" : "Mumbai",
			"state" : "MH",
			"country" : "India"
		} 
	}`

	var employee Employee
	err = json.Unmarshal([]byte(jsonData1), &employee)
	if err != nil {
		fmt.Println("error in unmarshalling : ", err)
	}
	fmt.Println(employee)

	listOfCityState := []Address{
		Address{City: "Mumbai", State: "MH"},
		Address{City: "Pune", State: "MH"},
		Address{City: "Indore", State: "MP"},
		Address{City: "Bangalore", State: "KT"},
	}

	fmt.Println(listOfCityState)
	jsonList, err := json.Marshal(&listOfCityState)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(string(jsonList))

	jsonData2 := `{"name" : "Aditya" , "age" : 30 , "address" : {"city" : "New York" , "state" : "NY"}}`
	var data map[string]interface{}
	err = json.Unmarshal([]byte(jsonData2), &data)
	if err != nil {
		fmt.Println("error in unmarshalling : ", err)
		return
	}
	fmt.Println(data["address"])

}

type Employee struct {
	FullName string  `json:"name"`
	ID       int     `json:"EmployeeId"`
	Address  Address `json:"address"`
}
