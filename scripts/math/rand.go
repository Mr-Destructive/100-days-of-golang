package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// random integer generation
	x := rand.Int()
	fmt.Println(x)

	// random number generation till range
	for i := 0; i < 5; i++ {
		y := rand.Intn(10)
		fmt.Println(y)
	}
}
