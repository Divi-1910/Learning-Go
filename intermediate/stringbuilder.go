package main

import (
	"fmt"
	"strings"
)

func main() {
	var builder strings.Builder

	builder.WriteString("Hello World")
	builder.WriteString(" ")
	builder.WriteString("!!")

	result := builder.String()
	fmt.Println(result)

	builder.WriteRune(' ')
	builder.WriteString("Hello World")
	result = builder.String()
	fmt.Println(result)

	builder.Reset()
	builder.WriteString("New Start")
	result = builder.String()
	fmt.Println(result)

}
