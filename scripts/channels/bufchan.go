package main

import (
	"fmt"
	"sync"
)

func main() {
	// Create a buffered channel with a capacity of 2
	buffchan := make(chan int, 2)

	// WaitGroup to wait for goroutines to finish
	wg := sync.WaitGroup{}
	wg.Add(2)

	// Start two goroutines
	for i := 1; i <= 2; i++ {
		go func(n int) {
			// Send values to the buffered channel
			buffchan <- n
			wg.Done()
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Close the buffered channel since we are done sending values
	close(buffchan)

	// Read values from the buffered channel
	for c := range buffchan {
		fmt.Println(c)
	}
}
