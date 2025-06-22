package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

/*
	Why use Context?
		Cancellation
		Timeouts
		Values

	Basic Concepts
		Context Creation
			context.Background()
			context.TODO()
		Context Hierarchy (how contexts are created and derived)
			context.WithCancel()
			context.WithDeadline()
 			context.WithTimeline()
			context.WithValue()

	Practical Usage
		Context Cancellation
		Timeouts and Deadlines
		Context values

	Best Practices
		Avoid Storing Contexts in Structs
		Propagating context correctly
		handling context values
		handling context cancellations
		avoid creating contexts in loops

	Common Pitfalls
		Ignoring Context cancellation
		Misusing context values

*/

func doWork(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Work Cancelled", ctx.Err())
		default:
			fmt.Println("Working...")
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func checkEvenOdd(ctx context.Context, num int) string {
	select {
	case <-ctx.Done():
		return "Operation Cancelled"
	default:
		if num%2 == 0 {
			return fmt.Sprintf("%d is even", num)
		} else {
			return fmt.Sprintf("%d is odd", num)
		}
	}
}

func main() {
	todoContext := context.TODO()
	BgContext := context.Background()

	ctx := context.WithValue(todoContext, "name", "Divyansh")
	fmt.Println(ctx)
	fmt.Println(ctx.Value("name"))

	ctx = context.WithValue(BgContext, "city", "Pune")
	fmt.Println(ctx)
	fmt.Println(ctx.Value("city"))

	fmt.Println()

	ctx = context.TODO()
	result := checkEvenOdd(ctx, 4)
	fmt.Println("Result with Context TODO : ", result)

	ctx = context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	time.Sleep(2 * time.Second)
	result = checkEvenOdd(ctx, 9)
	fmt.Println("Result with context Background : ", result)

	fmt.Println()
	fmt.Println()

	rootCtx := context.Background()
	rootCtx, cancel = context.WithTimeout(rootCtx, 2*time.Second)
	defer cancel()

	rootCtx = context.WithValue(rootCtx, "name", "Divyansh")
	go doWork(rootCtx)
	time.Sleep(3 * time.Second)

	name := rootCtx.Value("name")
	fmt.Println(name)

	ctx = context.Background()
	ctx, cancel = context.WithCancel(ctx)
	defer cancel()

	go func() {
		time.Sleep(2 * time.Second) // simulating a heavy task
		cancel()
	}()
	time.Sleep(3 * time.Second)

	ctx = context.WithValue(ctx, "RequestID", 1)
	logsWithContext(ctx, "This is a log message")

}

// contextual logging
func logsWithContext(ctx context.Context, message string) {
	requestIDVal := ctx.Value("RequestID")
	log.Printf("RequestID : %v - %v", requestIDVal, message)
}
