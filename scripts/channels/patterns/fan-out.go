package main

import (
	"fmt"
	"os"
	"sync"
)

func readFile(file string, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("Error reading from %s: %v\n", file, err)
		return
	}

	ch <- string(content)
}

func main() {
	files := []string{"data/f1.txt", "data/f2.txt"}

	var wg sync.WaitGroup
	ch := make(chan string)

	for _, f := range files {
		wg.Add(1)
		go readFile(f, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	var fileData []string
	for content := range ch {
		fileData = append(fileData, content)
	}

	fmt.Printf("Read %d files\n", len(fileData))
	fmt.Printf("Contents:\n%s", fileData)
}
