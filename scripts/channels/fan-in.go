package main

import (
	"fmt"
	"os"
	"sync"
)

func readFile(file string, ch chan<- string) {
	content, _ := os.ReadFile(file)
	fmt.Println("Reading from", file, "as :: ", string(content))
	ch <- string(content)
	close(ch)
}

func merge(chs ...<-chan string) string {
	var wg sync.WaitGroup
	out := ""

	for _, ch := range chs {
		wg.Add(1)
		go func(c <-chan string) {
			for s := range c {
				out += s
			}
			wg.Done()
		}(ch)
	}

	wg.Wait()
	return out
}

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go readFile("data/f1.txt", ch1)
	go readFile("data/f2.txt", ch2)

	merged := merge(ch1, ch2)

	fmt.Println(merged)
}
