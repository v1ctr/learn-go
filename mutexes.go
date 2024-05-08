package main

import (
	"fmt"
	"sync"
)

// Container holds a map of counters; since we want to update it concurrently from multiple goroutines,
// we add a Mutex to synchronize access.
// Note that mutexes must not be copied, so if this struct is passed around, it should be done by pointer.
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	// Lock the mutex before accessing counters; unlock it at the end of the function using a defer statement.
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	c := Container{

		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	wg.Wait()
	fmt.Println(c.counters)
}
