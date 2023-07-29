package main

import (
	"fmt"
	"sync"
	"time"
)

func process(pid int, wg *sync.WaitGroup) {
	fmt.Printf("Started process %d\n", pid)
	time.Sleep(1 * time.Second)
	fmt.Printf("Completed process %d\n", pid)
	defer wg.Done()
}

func main() {

	now := time.Now()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All processes completed")
	end := time.Now()
	fmt.Println(end.Sub(now))
}
