package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		ch <- 1
		time.Sleep(2 * time.Second)
		fmt.Println("2 Second GoRoutine Finished !! ")
	}()

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("3 Second GoRoutine Finished !! ")
	}()

	receiver := <-ch
	fmt.Println(receiver)
	fmt.Println("End of Program")

}
