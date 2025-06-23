package main

import (
	"fmt"
	"sync"
	"time"
)

type RateLimiter struct {
	mu        sync.Mutex
	count     int
	limit     int
	window    time.Duration
	resetTime time.Time
}

func NewFixedWindowAlgo(limit int, window time.Duration) *RateLimiter {
	rateLimiter := &RateLimiter{
		limit:  limit,
		window: window,
	}
	return rateLimiter
}

func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	now := time.Now()
	if now.After(rl.resetTime) {
		rl.resetTime = now.Add(rl.window)
		rl.count = 0
	}

	if rl.count < rl.limit {
		rl.count++
		return true
	}
	return false
}

func main() {
	rateLimiter := NewFixedWindowAlgo(5, 2*time.Second)
	var wg sync.WaitGroup
	for range 10 {
		wg.Add(1)
		go func() {
			if rateLimiter.Allow() {
				fmt.Println("Request Allowed")
			} else {
				fmt.Println("Request Denied")
			}
			time.Sleep(500 * time.Millisecond)
			wg.Done()
		}()
	}
	wg.Wait()

}
