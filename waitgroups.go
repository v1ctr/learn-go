package main

import (
	"fmt"
	"sync"
	"time"
)

// To wait for multiple goroutines to finish, we can use a wait group.

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	// This WaitGroup is used to wait for all the goroutines launched here to finish.
	// Note: if a WaitGroup is explicitly passed into functions, it should be done by pointer.
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		// Wrap the worker call in a closure that makes sure to tell the WaitGroup that this worker is done.
		// This way the worker itself does not have to be aware of the concurrency primitives involved in its execution.
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	// Block until the WaitGroup counter goes back to 0; all the workers notified they’re done.
	wg.Wait()

}
