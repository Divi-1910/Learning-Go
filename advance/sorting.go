package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByAge []Person

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

type By func(p1, p2 *Person) bool
type PersonSorter struct {
	people []Person
	by     func(p1, p2 *Person) bool
}

func (p *PersonSorter) Len() int {
	return len(p.people)
}

func (p *PersonSorter) Less(i, j int) bool {
	return p.by(&p.people[i], &p.people[j])
}
func (p *PersonSorter) Swap(i, j int) {
	p.people[i], p.people[j] = p.people[j], p.people[i]
}

func (by By) Sort(people []Person) {
	ps := &PersonSorter{
		people: people,
		by:     by,
	}
	sort.Sort(ps)
}

func main() {
	numbers := []int{5, 3, 4, 1, 2}

	sort.Ints(numbers)
	fmt.Println(numbers)

	stringSlice := []string{"Divyansh", "Sharvari", "Aditya", "Amruta"}
	sort.Strings(stringSlice)
	fmt.Println(stringSlice)

	people := []Person{
		{"Bob", 20},
		{"Alice", 15},
		{"Manna", 18},
	}
	sort.Sort(ByAge(people))
	fmt.Println(people)

	name := func(p1, p2 *Person) bool {
		return p1.Name < p2.Name
	}
	ageDesc := func(p1, p2 *Person) bool {
		return p1.Age > p2.Age
	}

	By(name).Sort(people)
	fmt.Println(people)

	By(ageDesc).Sort(people)
	fmt.Println(people)

	fruits := []string{"Apple", "Orange", "Pear", "Banana", "Grapes", "Cherry"}
	sort.Slice(fruits, func(i, j int) bool {
		ch1 := fruits[i][len(fruits[i])-1]
		ch2 := fruits[j][len(fruits[j])-1]
		return ch1 > ch2

	})
	fmt.Println(fruits)

}

/*
Why use sorting by functions
	Custom criteria
	Reusability
	Flexibility

Sorting by Functions
	`sort.Interface` consists of three methods
		Len() int
		Less(i , j int)  bool
		Swap(i , j int)


*/
