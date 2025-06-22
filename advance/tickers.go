package main

import (
	"fmt"
	"time"
)

/*
Why use Tickers?
	Consistency
	Simplicity

Best Practices for using tickers
	stop tickers when no longer needed
	avoid blocking operations
	handle ticker stopping gracefully

*/

func periodicTask() {
	fmt.Println("Periodic Task is performed at : ", time.Now().String())
}

func main() {
	ticker := time.NewTicker(1 * time.Second)
	i := 0
	for range ticker.C {
		i++
		fmt.Printf("Value of i at %d \n", i)
		if i > 5 {
			ticker.Stop()
			break
		}
	}

	// periodic task execution
	ticker = time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	//for {
	//	select {
	//	case <-ticker.C:
	//		periodicTask()
	//	}
	//}

	ticker = time.NewTicker(1 * time.Second)
	stop := time.After(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case tick := <-ticker.C:
			fmt.Println("Ticker is running at tick : ", tick)
		case <-stop:
			fmt.Println("Ticker is stopping ")
			return
		}
	}

}
