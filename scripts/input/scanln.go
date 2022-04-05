package main

import "fmt"

func main() {
	// Scanln for input only beofre line break
	var s string
	fmt.Println("Enter a string: ")
	fmt.Scanln(&s)
	fmt.Println("Inputted string : ", s)
}
