package main

import (
	"fmt"
	"time"
)

func main() {

	// Here we’ll use the select builtin on the channel to await the values as they arrive every 500ms.
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// Tickers can be stopped like timers.
	// Once a ticker is stopped it won’t receive any more values on its channel.
	// We’ll stop ours after 1600ms.
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
