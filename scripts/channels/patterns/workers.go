package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int) {
	for job := range jobs {
		fmt.Printf("Worker %d: Processing Job %d\n", id, job)
		time.Sleep(100 * time.Millisecond) // Simulating work
	}
}

func main() {
	const numWorkers = 5
	const numJobs = 100

	jobs := make(chan int, numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker(i, jobs)
	}

	for i := 0; i < numJobs; i++ {
		jobs <- i
	}

	close(jobs)

	time.Sleep(2 * time.Second)
}
