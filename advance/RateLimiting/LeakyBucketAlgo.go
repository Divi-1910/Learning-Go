package main

import (
	"fmt"
	"sync"
	"time"
)

type LeakyBucket struct {
	capacity     int
	leakRate     time.Duration
	tokens       int
	lastLeakTime time.Time
	mu           sync.Mutex
}

func NewLeakyBucket(capacity int, leakRate time.Duration) *LeakyBucket {
	return &LeakyBucket{
		capacity:     capacity,
		leakRate:     leakRate,
		tokens:       capacity,
		lastLeakTime: time.Now(),
	}
}

func (lb *LeakyBucket) Allow() bool {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	now := time.Now()
	elapsed := now.Sub(lb.lastLeakTime)
	tokensTodAdd := int(elapsed / (lb.leakRate))

	lb.tokens = lb.tokens + tokensTodAdd
	if lb.tokens >= lb.capacity {
		lb.tokens = lb.capacity
	}
	lb.lastLeakTime = lb.lastLeakTime.Add(time.Duration(tokensTodAdd) * lb.leakRate)
	if lb.tokens > 0 {
		lb.tokens--
		return true
	}
	return false

}

func main() {
	leakyBucketInst := NewLeakyBucket(5, 500*time.Millisecond)

	for range 10 {
		if leakyBucketInst.Allow() {
			fmt.Println("Leaky Bucket Allow")
		} else {
			fmt.Println("Leaky Bucket Deny")
		}
		time.Sleep(200 * time.Millisecond)
	}
}
