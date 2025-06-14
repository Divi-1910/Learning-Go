package main

import "fmt"

func swap[T any](x T, y T) (T, T) {
	return y, x
}

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) push(x T) {
	s.elements = append(s.elements, x)
}

func (s *Stack[T]) pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	x := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return x, true
}

func (s *Stack[T]) isEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack[T]) print() {
	if s.isEmpty() {
		fmt.Println("empty stack")
		return
	}
	for _, e := range s.elements {
		fmt.Printf("%v ", e)
	}
	fmt.Println()
}

func main() {
	x, y := 1, 2
	x, y = swap(x, y)
	fmt.Println(x, y)

	arr := []int{1, 2, 3}
	fmt.Println(arr)

	st := Stack[int]{}
	st.push(3)
	st.push(4)
	st.push(5)
	st.push(6)
	st.print()
	st.pop()
	st.print()
	st.isEmpty()

	strst := Stack[string]{}
	strst.push("hello")
	strst.push("world")
	strst.print()
	strst.pop()
	strst.print()
	strst.pop()
	strst.print()

}
