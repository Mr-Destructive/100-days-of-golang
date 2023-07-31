package main

import (
	"fmt"
	"sync"
	"time"
)

func reader(id int, count *int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	if !mutex.TryLock() {
		fmt.Printf("Reader %d blocked!\n", id)
		return
	}
	defer mutex.Unlock()
	fmt.Printf("Reader %d: read count %d\n", id, *count)

}
func writer(id, increment int, count *int, wg *sync.WaitGroup, mutex *sync.Mutex) {

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
	var mutex sync.Mutex
	readers := 5
	writers := 3

	wg.Add(readers)
	for i := 0; i < readers; i++ {
		go reader(i, &count, &wg, &mutex)
	}

	wg.Add(writers)
	for i := 0; i < writers; i++ {
		go writer(i, 1, &count, &wg, &mutex)
	}
	wg.Wait()

}
