package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Key concepts of Sync.RWMutex
	Read lock - RLock
	Write lock - Lock
	Unlock - Unlock and RUnlock

When to use RWMutex
	Read-Heavy Workloads
	Shared Data Structures

How RWMutex Works
	Read Lock Behaviour
	Write Lock Behaviour
	Lock Contention and Starvation

*/

var rwmutex sync.RWMutex
var counter int

func readCounter(wg *sync.WaitGroup) {
	defer wg.Done()
	rwmutex.RLock()
	fmt.Println("Read Counter : ", counter)
	rwmutex.RUnlock()
}

func writeCounter(wg *sync.WaitGroup, value int) {
	defer wg.Done()
	rwmutex.Lock()
	counter = value
	rwmutex.Unlock()
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go readCounter(&wg)
	}

	wg.Add(1)
	time.Sleep(time.Second)
	go writeCounter(&wg, 10)

	wg.Wait()

}
