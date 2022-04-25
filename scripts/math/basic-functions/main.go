package main

import (
	"fmt"
	"math"
)

func main() {

	//absolute function
	a := 45
	b := 100
	diff := a - b
	fmt.Println(diff)

	absolute_diff := math.Abs(float64(a) - float64(b))
	fmt.Println(absolute_diff)

	//min-max
	a = 120
	b = 54
	minimum := math.Min(float64(a), float64(b))
	maximum := math.Max(float64(a), float64(b))
	fmt.Printf("Min of %v and %v is %v \n", a, b, minimum)
	fmt.Printf("Max of %v and %v is %v \n", a, b, maximum)

	//pow and pow-10
	var x float64 = 3
	var y float64 = 4
	z := math.Pow(x, y)
	z10 := math.Pow10(int(x))
	fmt.Println("X ^ Y = ", z)
	fmt.Println("10 ^ X = ", z10)

	// sqrt and cbrt

	var k float64 = 125
	sqrt_of_k := math.Sqrt(k)
	cbrt_of_k := math.Cbrt(k)

	fmt.Printf("Square root of %v = %v \n", k, sqrt_of_k)
	fmt.Printf("Cube root of %v = %v \n", k, cbrt_of_k)

	// truncate

	var p float64 = 445.235
	trunc_p := math.Trunc(p)
	fmt.Printf("Truncated value of %v = %v \n", p, trunc_p)
	p = 123.678
	trunc_p = math.Trunc(p)
	fmt.Printf("Truncated value of %v = %v \n", p, trunc_p)

	// ceil

	var c float64 = 33.25
	ceil_c := math.Ceil(c)
	fmt.Printf("Ceiled value of %v = %v \n", c, ceil_c)
	c = 134.78
	ceil_c = math.Ceil(c)
	fmt.Printf("Ceiled value of %v = %v \n", c, ceil_c)
}
