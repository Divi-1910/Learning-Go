package main

import (
	"fmt"
	"time"
)

/*
Why use Stateful Goroutines ?
	State Management
	Concurrency
	Task Execution

Key Concepts of stateful Goroutines
	State Preservation
	Concurrency Management
	LifeCycle Management

*/

type StatefulWorker struct {
	count int
	ch    chan int
}

func (worker *StatefulWorker) Start() {
	go func() {
		for {
			select {
			case value := <-worker.ch:
				worker.count += value
				fmt.Println("Current Count : ", worker.count)
			}
		}
	}()
}

func (worker *StatefulWorker) Send(value int) {
	worker.ch <- value
}

func main() {
	stWorker := &StatefulWorker{
		ch: make(chan int),
	}
	stWorker.Start()

	for i := range 5 {
		stWorker.Send(i)
		time.Sleep(500 * time.Millisecond)
	}

}
