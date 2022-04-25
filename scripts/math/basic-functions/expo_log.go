package main

import (
	"fmt"
	"math"
)

func main() {
	// exponential function
	var x float64 = 2
	y := math.Exp(x)
	fmt.Println("e^x = ", y)
	var n float64 = 3.5
	y = math.Exp2(n)
	fmt.Println("2^n = ", y)

	// Logarithmic function
	y = math.Log(x)
	fmt.Println("natural log x = ", y)

	n = 128
	y = math.Log2(n)
	fmt.Println("Log2 of 100 = ", y)
}
