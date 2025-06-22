package main

import (
	"fmt"
	"sync"
)

/*
Why use Mutexes?
	Data Integrity
	Synchronization
	Avoid Race Conditions

Basic Operations
	Lock()
	Unlock()
	TryLock()

Mutex and Performance:
	Contention
	Granularity

Why Mutual Exclusion is Important
	Data Integrity
	Consistency
	Safety

How Mutual Exclusion is Achieved
	Locks (Mutexes)
	Semaphores
	Monitors
	Critical Sections

Why they are often used in structs
	Encapsulation
	Convenience
	Readability

*/

type counter struct {
	mu    sync.Mutex
	count int
}

func (c *counter) increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++

}

func (c *counter) decrement() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count--
}

func (c *counter) getValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {
	var wg sync.WaitGroup
	cn := &counter{} // allocates memory to a pointer instance and returns the pointer to it

	numGoroutines := 10

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				cn.increment()
			}
		}()
	}

	wg.Wait()
	fmt.Println("Final Counter value is : ", cn.getValue())

	fmt.Println()
	fmt.Println()

	var cnt int
	var newwg sync.WaitGroup
	var mu sync.Mutex

	numGoroutines = 5
	newwg.Add(numGoroutines)

	increment := func() {
		defer newwg.Done()
		for i := 0; i < 1000; i++ {
			mu.Lock()
			cnt++
			mu.Unlock()
		}
	}

	for i := 0; i < numGoroutines; i++ {
		go increment()
	}
	newwg.Wait()
	fmt.Println("Final Counter value is : ", cnt)

}

/*
Best Practices for using mutexes
	Minimize lock duration
	Avoid Nested Locks
	Prefer sync.RWMutex for read-heavy workloads
	check for deadlocks
	use `defer` for unlocking

Common Pitfalls
	Deadlocks
	Performance issues
	Starvation

*/
