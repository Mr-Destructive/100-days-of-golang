package main

import (
	"fmt"
	"time"
)

func process() {
	time.Sleep(time.Second)
	fmt.Println("Hello World!")
}

func main() {
	start := time.Now()
	go process()
	go process()
	end := time.Now()
	duration := end.Sub(start)
	time.Sleep(time.Second)
	fmt.Println(duration)
}
