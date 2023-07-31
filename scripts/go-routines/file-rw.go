package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

func reader(id int, filename string, wg *sync.WaitGroup, mutex *sync.RWMutex) {
	defer wg.Done()
	if !mutex.TryRLock() {
		fmt.Printf("Reader %d blocked!\n", id)
		return
	}
	defer mutex.RUnlock()
	f, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Reader %d read: %s\n", id, string(bytes))
}

func writer(id int, filename, content string, wg *sync.WaitGroup, mutex *sync.RWMutex) {

	defer wg.Done()
	mutex.Lock()
	defer mutex.Unlock()
	f, err := os.OpenFile(filename, os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	f.WriteString(content)
	fmt.Printf("Writer %d: wrote -> %s\n", id, content)
}

func main() {

	var wg sync.WaitGroup
	var mutex sync.RWMutex
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go reader(i, "data/f2.txt", &wg, &mutex)
	}

	wg.Add(2)
	for i := 0; i < 2; i++ {
		content := fmt.Sprintf("line %d.\n", i)
		go writer(i, "data/f2.txt", content, &wg, &mutex)
	}

	wg.Wait()

}
