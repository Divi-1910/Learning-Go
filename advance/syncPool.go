package main

import (
	"fmt"
	"sync"
)

/*

Sync.Pool -> maintaining collection of reusable objects

Key Concepts of sync.Pool
	Object Pooling
	Object Retrieval and Return

Methods of sync.Pool
	Get()
	Put(interface{})
	New (Optional)

*/

type Person struct {
	Name string
	Age  int
}

func main() {
	var pool = sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating a new Person")
			return &Person{}
		},
	}

	// Get an Object from the pool
	person1 := pool.Get().(*Person)
	person1.Name = "Jack"
	person1.Age = 18
	fmt.Println("person1 : ", person1)

	pool.Put(person1)
	fmt.Println("Returned the person to pool")

	person2 := pool.Get().(*Person)
	fmt.Println("person2 : ", person2)

	person3 := pool.Get().(*Person)
	fmt.Println("Got Person Again", person3)
	person3.Name = "Jane"

	pool.Put(person2)
	pool.Put(person3)

	person4 := pool.Get().(*Person)
	fmt.Println("person4 : ", person4)

	person5 := pool.Get().(*Person)
	fmt.Println("person5 : ", person5)

}
