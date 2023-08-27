package main

import (
	"fmt"
	"sync"
)

func sendMessage(ch chan string, message string, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- message
}

func main() {
	var wg sync.WaitGroup

	ch1 := make(chan string, 2)
	ch2 := make(chan string, 2)
	wg.Add(2)

	go sendMessage(ch1, "Hello, Gophers!", &wg)
	go sendMessage(ch2, "Hello, Hamsters!", &wg)

	go func() {
		defer wg.Done()
		wg.Wait()
		close(ch1)
		close(ch2)
	}()
	ch1 <- "new message to c1"
	ch2 <- "new message to c2"

	select {
	case <-ch1:
		fmt.Println("Received from ch1")
	case ch1 <- "new message to c1":
		fmt.Println("Sent to ch1")
	case <-ch2:
		fmt.Println("Received from ch2")
	case ch2 <- "new message to c2":
		fmt.Println("Sent to ch2")
	}

}
