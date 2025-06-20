package main

import "fmt"

func main() {
	var a int = 42
	b := int64(42)
	c := float64(b)

	d := 3.14
	f := int(d)
	fmt.Println(a, b, c, d, f)

	g := "Hello"
	h := []byte(g)
	fmt.Println(string(h), h)

	i := []int{67, 72}
	fmt.Println(i)

}
