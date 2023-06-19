package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()
	time.Sleep(time.Second * 3)
	t2 := time.Now()
	duration := t2.Sub(t1)
	fmt.Println(duration)

	ticker := time.NewTicker(time.Second * 2)
	counter := 0
	for {
		select {
		case <-ticker.C:
			// api calls, call to database after specific time intervals
			counter++
			fmt.Println("Tick", counter)
		}
		if counter == 5 {
			ticker.Stop()
			return
		}
	}
}
