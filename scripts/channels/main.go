package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	defer close(ch)

	go func() {
		ch <- 1
	}()

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
