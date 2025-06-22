package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Why use Wait Groups?
	Synchronization
	Co-ordination
	Resource Management

Basic Operations
	Add (delta int)
	Done()
	Wait()
*/

// basic examples without using channels
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Worker ", id, "starting...")
	time.Sleep(1 * time.Second)
	fmt.Println("Worker ", id, "finished")
}

func worker1(id int, results chan<- int, tasks <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker1ID %d starts working...\n", id)
	time.Sleep(1 * time.Second) // simulating work
	for {
		t, ok := <-tasks
		if !ok {
			fmt.Println("Tasks channel is closed.")
			break
		}
		results <- t * 2
	}
	fmt.Printf("Worker1ID %d finished working...\n", id)
}

func main() {
	var wg sync.WaitGroup
	numWorkers := 3
	wg.Add(numWorkers) // add worker count in wait group

	for i := 0; i < numWorkers; i++ {
		go worker(i, &wg) // launch workers
	}
	wg.Wait() // blocking mechanism
	fmt.Println("All workers finished")

	fmt.Println()
	fmt.Println()

	var wg1 sync.WaitGroup
	numWorkers = 3
	numJobs := 3
	results := make(chan int, numJobs)
	tasks := make(chan int, numJobs)

	wg1.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker1(i, results, tasks, &wg1)
	}

	for i := 0; i < numJobs; i++ {
		tasks <- i
	}
	close(tasks)

	go func() {
		wg1.Wait()
		close(results)
	}()

	for {
		r, ok := <-results
		if !ok {
			fmt.Println("Channel closed")
			break
		}
		fmt.Println(r)
	}

	var wg sync.WaitGroup
	tasks := []string{"digging", "laying bricks", "painting walls"}

	for i, task := range tasks {
		wk := Worker{
			ID:   i,
			Task: task,
		}
		wg.Add(1)
		go wk.PerformTask(&wg)
	}

	wg.Wait()

	fmt.Println("Construction is finished")

}

type Worker struct {
	ID   int
	Task string
}

func (w *Worker) PerformTask(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Worker ", w.ID, "starting...", w.Task)
	time.Sleep(1 * time.Second)
	fmt.Println("Worker ", w.ID, "finished the Task ", w.Task)
}

/*
Best Practices
	Avoid Blocking calls inside goroutines
	use defer to call done
	ensure proper use of add
	handle large numbers of goroutines
	use defer for unlocking

common Pitfalls
	mismatch between add and done
	avoid creating deadlocks

*/
