package main

import (
	"fmt"
	"math"
	"reflect"
)

type Geometry interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func (r Rect) Area() float64 {
	return r.width * r.height
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rect) Perimeter() float64 {
	perim := 2 * (r.width + r.height)
	return perim
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c Circle) Diameter() float64 {
	return 2 * c.radius
}

func measure(g Geometry) {
	fmt.Println(g)
	fmt.Println(g.Area())
	fmt.Println(g.Perimeter())
}

// capitalized method names are exported and accessible outisde the package
// uncapitalized ones are private to the package

func Printer(val ...any) {
	for _, v := range val {
		fmt.Println(v, " - ", reflect.TypeOf(v))
	}
}

func TypePrinter(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float64")
	case string:
		fmt.Println("string")
	}
}

func main() {
	r := Rect{width: 10, height: 5}
	c := Circle{radius: 7}

	measure(r)
	measure(c)

	Printer(1, 2, "Hello", 2.343)

}
