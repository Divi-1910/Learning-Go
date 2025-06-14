package main

import (
	"fmt"
	"maps"
)

func main() {

	// var mapVariableName map[keyType]valueType
	// mapVariable = make(map[keyType]valueType
	// mapVariable = map[keyType]valueType{
	//		key1: value1,
	//		key2: value2
	//	}

	myMap := make(map[string]int)

	fmt.Println(myMap)

	myMap["key"] = 1
	myMap["key2"] = 2

	fmt.Println(myMap)
	fmt.Println(myMap["key3"])
	myMap["key25"] = 4

	fmt.Println(myMap)
	_, exists := myMap["key3"]
	fmt.Println(exists)

	if maps.Equal(myMap, myMap) {
		fmt.Println("these maps are equal")
	}

	for k, v := range myMap {
		fmt.Println(k, v)
	}

	var Mymap3 map[string]string

	if Mymap3 == nil {
		fmt.Println("mymap3 is nil")
	}

	val := Mymap3["0"]
	fmt.Println(val)

}
