package main

import (
	"fmt"
	"sync"
	"time"
)

func reader(id int, count *int, wg *sync.WaitGroup, mutex *sync.RWMutex) {
	defer wg.Done()
	if !mutex.TryRLock() {
		fmt.Printf("Reader %d blocked!\n", id)
		return
	}
	defer mutex.RUnlock()
	fmt.Printf("Reader %d: read count %d\n", id, *count)

}
func writer(id, increment int, count *int, wg *sync.WaitGroup, mutex *sync.RWMutex) {

	defer wg.Done()
	mutex.Lock()
	defer mutex.Unlock()
	*count += increment
	time.Sleep(5 * time.Millisecond)
	fmt.Printf("Writer %d: wrote count %d\n", id, *count)
}

func main() {

	count := 1
	var wg sync.WaitGroup
	var mutex sync.RWMutex
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go reader(i, &count, &wg, &mutex)
	}

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go writer(i, 1, &count, &wg, &mutex)
	}
	wg.Wait()

}
