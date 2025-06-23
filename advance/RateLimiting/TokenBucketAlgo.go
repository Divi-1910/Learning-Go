package main

import (
	"fmt"
	"time"
)

type RateLimiter struct {
	tokens     chan struct{}
	refillTime time.Duration
}

func NewRateLimiter(rateLimit int, refillTime time.Duration) *RateLimiter {
	r1 := &RateLimiter{
		tokens:     make(chan struct{}, rateLimit),
		refillTime: refillTime,
	}
	for range rateLimit {
		r1.tokens <- struct{}{}
	}
	go r1.startRefill()
	return r1
}

func (r1 *RateLimiter) startRefill() {
	ticker := time.NewTicker(r1.refillTime)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			select {
			case r1.tokens <- struct{}{}:
			default:
			}
		}
	}
}

func (r1 *RateLimiter) Allow() bool {
	select {
	case <-r1.tokens:
		return true
	default:
		return false
	}
}

func main() {
	rateLimiter := NewRateLimiter(5, 5*time.Second)
	for range 15 {
		if rateLimiter.Allow() {
			fmt.Println("Request Allowed")
		} else {
			fmt.Println("Request Denied")
		}
		time.Sleep(500 * time.Millisecond)
	}
}
