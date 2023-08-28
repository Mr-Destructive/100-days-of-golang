package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int, 5)
	go func() {
		for i := 0; i < 100; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	throttle := time.Tick(500 * time.Millisecond)

	for job := range jobs {
		<-throttle // limits to 5 jobs/sec
		go func(job int) {
			fmt.Println("Job:", job)
		}(job)
	}
}
