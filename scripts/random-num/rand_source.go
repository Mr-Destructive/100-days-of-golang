package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	rand_source := rand.New(source)
	for i := 0; i < 5; i++ {
		rand_num := rand_source.Int()
		fmt.Println(rand_num)
	}
}
