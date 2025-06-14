package main

import "fmt"

type Rectangle struct {
	length float64
	width  float64
}

func (r *Rectangle) Area() float64 {
	return r.width * r.length
}

func (r *Rectangle) updateRectangle(newlength float64, newwidth float64) {
	r.length = newlength
	r.width = newwidth
}

type Length int

func (l Length) isLargeLength() bool {
	return l > 100
}

func (Length) welcomeMessage() string {
	return "Hello and Welcome"
}

func main() {
	rect := Rectangle{length: 15, width: 10}
	fmt.Println("Area of Rectangle = ", rect.Area())

	len1 := Length(1000)
	fmt.Println(len1.isLargeLength())

	fmt.Println(len1.welcomeMessage())

	rect.updateRectangle(15, 15)
	fmt.Println("Area of New Rectangle = ", rect.Area())

}
