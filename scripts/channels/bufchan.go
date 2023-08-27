package main

import (
	"fmt"
)

func main() {
	buffchan := make(chan int, 2)

	buffchan <- 1
	buffchan <- 2

	fmt.Println(<-buffchan)
	fmt.Println(<-buffchan)
}
