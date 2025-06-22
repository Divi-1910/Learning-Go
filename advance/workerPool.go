package main

import (
	"fmt"
	"time"
)

/*
Why use Worker Pools?
	Resource Management
	Task Distribution
	Scalability

Conceptual Model
	Tasks
	Workers
	Task Queue

Implementation Steps
	Create a task channel
	create worker goroutines
	Distribute tasks
	Graceful shutdowns

Advanced Worker pool Patterns
	dynamic worker pools
	task prioritization
	error handling
	worker pool with task prioritization

Best Practices for worker pools
	limit the number of workers
	handle worker lifecycle
	implement timeouts
	monitor and log

*/

func worker(id int, tasks <-chan int, results chan<- int) {
	for task := range tasks {
		fmt.Println("Worker ", id, " starts task ", task)
		time.Sleep(1 * time.Second) // simulating work
		results <- task * 2
	}
}

// BASIC WORKER POOL PATTERN
//func main() {
//	numWorkers := 3
//	numJobs := 10
//	tasks := make(chan int, numJobs)
//	results := make(chan int, numJobs)
//
//	// Create workers
//	for i := range numWorkers {
//		go worker(i, tasks, results)
//	}
//
//	// send values to the task channels
//	for i := range numJobs {
//		tasks <- i
//	}
//
//	close(tasks)
//
//	// collect the results
//	for range numJobs {
//		result := <-results
//		fmt.Println("Result : ", result)
//	}
//
//}

type ticketRequest struct {
	personID   int
	numTickets int
	cost       int
}

// simulate processing of ticket requests
func ticketProcessor(requests <-chan ticketRequest, results chan<- int) {
	for req := range requests {
		fmt.Printf("Processing %d ticket(s) of personId %d with total cost of %d\n", req.numTickets, req.personID, req.cost)
		time.Sleep(1 * time.Second) // simulate processing time
		results <- req.personID
	}
}

func main() {
	numRequests := 5
	price := 5
	ticketRequests := make(chan ticketRequest, numRequests)
	ticketResults := make(chan int, numRequests)

	// start the ticket processor
	for range 3 {
		go ticketProcessor(ticketRequests, ticketResults)
	}

	for i := range numRequests {
		ticketRequests <- ticketRequest{personID: i + 1, numTickets: (i + 1) * 2, cost: (i + 2) * price}
	}
	close(ticketRequests)
	for range numRequests {
		fmt.Printf("Ticket for personID %d processed successfully\n", <-ticketResults)
	}

}
