package main

import (
	"fmt"
	"sync"
)

func generate(nums []int, out chan<- int, wg *sync.WaitGroup) {
	fmt.Println("Stage 1")
	for _, n := range nums {
		fmt.Println("Number:", n)
		out <- n
	}
	close(out)
	wg.Done()
}

func square(in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	fmt.Println("Stage 2")
	for n := range in {
		sq := n * n
		fmt.Println("Square:", sq)
		out <- sq
	}
	close(out)
	wg.Done()
}

func print(in <-chan int, wg *sync.WaitGroup) {
	for n := range in {
		fmt.Println(n)
	}
	wg.Done()
}

func main() {
	input := []int{1, 2, 3, 4, 5}

	var wg sync.WaitGroup
	wg.Add(3)

	stage1 := make(chan int)
	stage2 := make(chan int)

	go generate(input, stage1, &wg)

	go square(stage1, stage2, &wg)

	go print(stage2, &wg)

	wg.Wait()
}
