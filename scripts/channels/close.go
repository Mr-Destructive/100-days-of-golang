package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 50; i++ {
			ch <- i
		}
		close(ch)
	}()

	for num := range ch {
		fmt.Println("Received:", num)
	}
}
