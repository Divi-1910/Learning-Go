package main

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
	Email   string   `xml:"email"`
	Address Address  `xml:"address"`
}

type Address struct {
	City    string `xml:"city"`
	State   string `xml:"state"`
	Country string `xml:"country"`
}

type Book struct {
	XMLName xml.Name `xml:"book"`
	Title   string   `xml:"title,attr"`
	Author  string   `xml:"author,attr"`
	ISBN    string   `xml:"isbn,attr"`
	Psuedo  string   `xml:"psuedo"`
}

func main() {
	person := Person{Name: "John", Age: 42, Email: "asdfgh@sdfgh.com", Address: Address{City: "New York", State: "NY", Country: "USA"}}

	xmldata, err := xml.Marshal(&person)
	if err != nil {
		fmt.Println("Error in Marshalling Data : ", err)
		return
	}
	fmt.Println("XML Data : ", string(xmldata))

	xmldata1, err := xml.MarshalIndent(&person, "", "  ")
	if err != nil {
		fmt.Println("Error in Marshalling Data : ", err)
		return
	}
	fmt.Println("XML Data : \n", string(xmldata1))

	xmlRaw := `<person>
	<name>John</name>
	<age>42</age>
	<email>asdfgh@sdfgh.com</email>
	<city>New York</city>
	<address>
	<city>New York</city>
	<state>NY</state>
	<country>USA</country>
	</address>
	</person>`

	var personalData Person
	err = xml.Unmarshal([]byte(xmlRaw), &personalData)
	if err != nil {
		fmt.Println("Error in Unmarshalling Data : ", err)
		return
	}
	fmt.Println(personalData)

	book := Book{
		Title:  "Go Programming Language",
		Author: "John Doe",
		ISBN:   "123-456-789-000",
		Psuedo: "psuedo",
	}

	xmlBookData, err := xml.MarshalIndent(&book, "", "  ")

	if err != nil {
		fmt.Println("Error in Marshalling Data : ", err)
		return
	}

	fmt.Println("XML Data : \n", string(xmlBookData))

}

/*
 Best Practices :
	1. Use XML tags
	2. Validate XML
	3. Handle Nested Structures
	4. Handle Errors
	5. Custom Marshalling/UnMarshalling

Real World Scenarios
	1. Web Services and APIs
	2. Spring Framework
	3. Microsoft .Net Applications

Data Interchange and Storage
	1. RSS and Atom Feeds
	2. Electronic Data Interchange (EDI)

Industry Standards
	1. HealthCare (HL7)
	2. Finance (FIXML)

*/
