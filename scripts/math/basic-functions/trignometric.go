package main

import (
	"fmt"
	"math"
)

func main() {
	// basic trigonometric functions
	var x float64 = math.Pi / 2
	sinx := math.Sin(x)
	cosx := math.Cos(x)
	tanx := math.Tan(x)
	fmt.Printf("Sin(%v) = %v \n", x, sinx)
	fmt.Printf("Cos(%v) = %v \n", x, cosx)
	fmt.Printf("Tan(%v) = %v \n", x, tanx)

	// hyperbolic trigonometric functions
	var h float64 = math.Pi / 2
	sinh := math.Sinh(h)
	cosh := math.Cosh(h)
	tanh := math.Tanh(h)
	fmt.Printf("Sinh(%v) = %v \n", h, sinh)
	fmt.Printf("Cosh(%v) = %v \n", h, cosh)
	fmt.Printf("Tanh(%v) = %v \n", h, tanh)

	// Inverse Trigonometric functions
	var y float64 = -1
	arc_sin := math.Asin(y)
	arc_cos := math.Acos(y)
	arc_tan := math.Atan(y)
	fmt.Printf("Sin^-1(%v) = %v \n", y, arc_sin)
	fmt.Printf("Cos^-1(%v) = %v \n", y, arc_cos)
	fmt.Printf("Tan^-1(%v) = %v \n", y, arc_tan)
}
