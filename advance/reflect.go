package main

import (
	"fmt"
	"reflect"
)

/*
Why use Reflection
	Dynamic type inspection
	generic programming
	serialization/deserialization

Few Methods
	reflect.TypeOf
	reflect.Value
	reflect.Valueof
	reflect.ValueOf().Elem()

*/

func main() {
	x := 42
	v := reflect.ValueOf(x)

	t := v.Type()
	fmt.Println("Type : ", t)
	fmt.Println("Value : ", v)
	fmt.Println("Kind : ", t.Kind())
	fmt.Println("Is Zero : ", v.IsZero())
	fmt.Println("Is String : ", t.Kind() == reflect.String)
	fmt.Println("Is Int : ", t.Kind() == reflect.Int)

	y := 10
	v1 := reflect.ValueOf(&y).Elem()
	fmt.Println("Value of v1 : ", v1)
	v2 := reflect.ValueOf(&y)
	fmt.Println("Value of v2 : ", v2.Type())

	fmt.Println("Original Type of v1 : ", v1.Int())
	v1.SetInt(18) // modifying values at runtime
	fmt.Println("Modified Value of v1 : ", v1.Int())

	var itf interface{} = "Hello"
	v3 := reflect.ValueOf(itf)

	fmt.Println("v3 Type : ", v3.Type())

	if v3.Kind() == reflect.String {
		fmt.Println("String : ", v3.String())
	}

	fmt.Println()
	fmt.Println()

	p := person{Name: "Alice", Age: 18}
	pv := reflect.ValueOf(p)
	pp := reflect.ValueOf(&p).Elem()

	for i := 0; i < pv.NumField(); i++ {
		fmt.Printf("Field %d : %v\n", i, pv.Field(i))
	}

	nameField := pp.FieldByName("Name")
	if nameField.CanSet() {
		nameField.SetString("Divyansh")
	}
	// Note - Make sure to respect the Go's visibility rules
	fmt.Println("Modified Person : ", p)

	fmt.Println()
	fmt.Println()

	g := Greeter{}
	t = reflect.TypeOf(g)
	v = reflect.ValueOf(g)
	fmt.Println("Type of Greeter : ", t)
	fmt.Println("Value of Greeter : ", v)
	var method reflect.Method

	for i := 0; i < t.NumMethod(); i++ {
		method = t.Method(i)
		fmt.Printf("Method %d : %s\n", i, method.Name)
	}

	m := v.MethodByName("Greet")
	results := m.Call([]reflect.Value{reflect.ValueOf("Divi")})

	fmt.Println(results[0])

}

type Greeter struct{}

func (g Greeter) Greet(name string) string {
	return "Hello " + name
}

type person struct {
	Name string
	Age  int
}
