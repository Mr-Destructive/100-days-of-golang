package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	defer close(ch)

	go func() {
		ch <- "Hello, Gophers!"
	}()

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// uncommenting the above will result in a deadlock
}
