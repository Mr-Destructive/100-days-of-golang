package main

import (
	"fmt"
	"sync"
)

func receiver(ch <-chan int, wg *sync.WaitGroup) {
	for i := range ch {
		fmt.Println("Received:", i)
	}
	wg.Done()
}

func sender(ch chan<- int, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("Sent:", i)
	}
	close(ch)
	wg.Done()
}

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go receiver(ch, &wg)
	go sender(ch, &wg)
	wg.Wait()
}
