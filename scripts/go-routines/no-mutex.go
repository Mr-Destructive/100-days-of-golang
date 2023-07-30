package main

import (
	"fmt"
	"os"
	"sync"
)

func WriteToFile(filename string, contents string, buffer *[]byte, wg *sync.WaitGroup) {
	defer wg.Done()
	contentsBytes := []byte(contents)
	*buffer = append(*buffer, contentsBytes...)
	err := os.WriteFile(filename, contentsBytes, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	var wg sync.WaitGroup
	var sharedBuffer []byte

	wg.Add(2)
	go WriteToFile("data/f1.txt", "Hello\n", &sharedBuffer, &wg)
	go WriteToFile("data/f1.txt", "World!\n", &sharedBuffer, &wg)
	wg.Wait()

	fmt.Println(string(sharedBuffer))
}
