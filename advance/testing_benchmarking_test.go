package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func Add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	result := Add(2, 2)
	expect := 4
	if result != expect {
		t.Errorf("Test Failed :  expect %d, but got %d", expect, result)
	}
}

// table driven tests
func TestAddTableDriven(t *testing.T) {
	test := []struct{ a, b, expected int }{
		{2, 3, 5},
		{1, 0, 1},
		{0, 1, 1},
		{1, 1, 2},
		{1, 2, 3},
	}

	for _, test := range test {
		result := Add(test.a, test.b)
		if result != test.expected {
			t.Errorf("Test Failed :  expect %d, but got %d", test.expected, result)
		}
	}

}

func BenchmarkAddSmallInput(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(2, 2)
	}
}

func BenchmarkAddLargeInput(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(100000000000000, 20000000000000)
	}
}

func GenerateRandomSlice(size int) []int {
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Int()
	}
	return slice
}

func SumSlice(slice []int) int {
	sum := 0
	for _, n := range slice {
		sum += n
	}
	return sum
}

func BenchmarkGenerateRandomSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateRandomSlice(1000)
	}
}

func BenchmarkSumSlice(b *testing.B) {
	slice := GenerateRandomSlice(1000)
	b.ResetTimer() // reset timer for not considering setup time
	for i := 0; i < b.N; i++ {
		SumSlice(slice)
	}
}

func TestGenerateRandomSlice(t *testing.T) {
	size := 1000
	slice := GenerateRandomSlice(size)
	if len(slice) != size {
		t.Errorf("Test Failed :  expected Slice Size %d, but got %d", size, len(slice))
	}
}

func TestAddSubtests(t *testing.T) {
	test := []struct{ a, b, expected int }{
		{2, 3, 5},
		{1, 0, 1},
		{0, 1, 1},
		{1, 1, 2},
		{1, 2, 3},
	}

	for _, test := range test {
		t.Run(fmt.Sprintf("TestAddSubtests Add(%d , %d) : ", test.a, test.b), func(t *testing.T) {
			result := Add(test.a, test.b)
			if result != test.expected {
				t.Errorf("SubTest Failed :  expect %d, but got %d", test.expected, result)
			}
		})
	}

}
