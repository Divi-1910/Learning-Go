package main

import (
	"fmt"
	"sync"
)

var (
	once sync.Once
)

func initialize() {
	fmt.Println("initialize , only once ")
}

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Goroutine ", i)
			// initialize()
			once.Do(initialize)
		}()
	}

	wg.Wait()

}
