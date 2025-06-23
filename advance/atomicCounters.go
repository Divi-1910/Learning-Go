package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
Why use atomic counters?
	performance
	simplicity
	Concurrency

Atomic Operations in go
	window duration
	Request Counting
	Reset Mechanism

sync/atomic package:
	atomic.AddInt32/atomic.AddInt64
	atomic.LoadInt32/atomic.LoadInt64
	atomic.StoreInt32/atomic.StoreInt64
	atomic.CompareAndSwap32/atomic.CompareAndSwapInt64

what does atomic mean?
	indivisible
	uninterruptible

Why use Atomic Operations
	Lost Updates
	Inconsistent Reads

How Atomic Operations work?
	lock-free
	memory-visibility

Issues without Atomic Operations
	Data Race
	Inconsistent Results

*/

type AtomicCounter struct {
	count int64
}

func (ac *AtomicCounter) Increment() {
	atomic.AddInt64(&ac.count, 1)
}

func (ac *AtomicCounter) Decrement() {
	atomic.AddInt64(&ac.count, -1)
}

func (ac *AtomicCounter) GetCount() int64 {
	return atomic.LoadInt64(&ac.count)
}

func main() {
	var wg sync.WaitGroup
	numGoroutines := 10
	counter := &AtomicCounter{}
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				counter.Increment()
			}
		}()
	}

	wg.Wait()
	fmt.Println("Final Counter value is : ", counter.GetCount())
}
