package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Key Concepts for sync.NewCond
	Condition Variables
	Mutex and Condition Variables

Methods of Sync.Cond
	Wait()
	Signal()
	Broadcast()

Example Usage of Sync.NewCond
	Producer-Consumer Example

*/

const BUFFERSIZE = 5

type Buffer struct {
	items []int
	mutex sync.Mutex
	cond  *sync.Cond
}

func NewBuffer() *Buffer {
	b := &Buffer{
		items: make([]int, 0, BUFFERSIZE),
	}
	b.cond = sync.NewCond(&b.mutex)
	return b
}

func (b *Buffer) Produce(item int) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	if len(b.items) == BUFFERSIZE {
		b.cond.Wait()
	}
	b.items = append(b.items, item)
	fmt.Println("Produced : ", item)
	b.cond.Signal()
}

func (b *Buffer) Consume() int {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if len(b.items) == 0 {
		b.cond.Wait()
	}

	item := b.items[0]
	b.items = b.items[1:]
	fmt.Println("Consumed : ", item)
	b.cond.Signal()
	return item
}

func Producer(buffer *Buffer, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		buffer.Produce(i + 1)
		time.Sleep(time.Millisecond * 100)
	}
}

func Consumer(buffer *Buffer, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		buffer.Consume()
		time.Sleep(time.Millisecond * 200)
	}
}

func main() {
	b := NewBuffer()
	var wg sync.WaitGroup

	wg.Add(2)
	go Producer(b, &wg)
	go Consumer(b, &wg)

	wg.Wait()
}
