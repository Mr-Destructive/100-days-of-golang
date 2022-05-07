package main

import "fmt"

func main() {
	x := 3
	y := 5
	// 3 -> 011
	// 5 -> 101
	fmt.Println("X AND Y = ", x&y)
	fmt.Println("X OR Y = ", x|y)
	fmt.Println("X EXOR Y = ", x^y)
	fmt.Println("X Right Shift 1  = ", x>>1)
	fmt.Println("X Right Shift 2  = ", x>>2)
	fmt.Println("Y Left Shift 1 = ", y<<1)
	fmt.Println("X  = ", x&^y)
}
