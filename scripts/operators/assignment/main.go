package main

import "fmt"

func main() {
	var a int = 100
	b := 20
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	a += 30
	fmt.Println("a = ", a)
	b -= 5
	fmt.Println("b = ", b)
	a *= b
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	a /= b
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	a %= b
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
}
