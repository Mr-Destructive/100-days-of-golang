package main

import "fmt"

func square(x int) int {
	fmt.Println(x * x)
	return x * x
}

func main() {

	// simple defer example
	fmt.Println("First")
	defer fmt.Println("Second Ahhh..")
	fmt.Println("Third")

	// multiple defer statements
	fmt.Println("bag")
	defer fmt.Println("book")
	defer fmt.Println("cap")
	fmt.Println("laptop")
	defer fmt.Println("wallet")
	fmt.Println("headphones")

	// defer keyword in function
	defer square(12)
	defer square(10)
	square(15)
}
