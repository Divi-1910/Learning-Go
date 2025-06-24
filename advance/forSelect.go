package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	quit := make(chan string)

	go func() {
		time.Sleep(5 * time.Second)
		close(quit)
	}()

	for {
		select {
		case t := <-ticker.C:
			fmt.Println("Tick at", t)
		case <-quit:
			fmt.Println("Quit")
			ticker.Stop()
			return
		}
	}
	
}
