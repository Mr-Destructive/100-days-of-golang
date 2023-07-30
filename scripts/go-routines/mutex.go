package main

import (
	"fmt"
	"os"
	"sync"
)

func WriteToFile(filename string, contents string, buffer *[]byte, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	contentBytes := []byte(contents)

	mutex.Lock()
	*buffer = append(*buffer, contentBytes...)

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	_, err = f.Write(contentBytes)
	if err != nil {
		fmt.Println(err)
	}
	mutex.Unlock()
}

func main() {
	var wg sync.WaitGroup
	var mut sync.Mutex
	var sharedBuffer []byte

	wg.Add(2)
	go WriteToFile("data/f1.txt", "Hello Gophers!\n", &sharedBuffer, &wg, &mut)
	go WriteToFile("data/f1.txt", "Welcome to Goworld!\n", &sharedBuffer, &wg, &mut)
	wg.Wait()

	fmt.Println(string(sharedBuffer))
}
