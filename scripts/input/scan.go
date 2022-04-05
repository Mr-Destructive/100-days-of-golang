package main

import "fmt"

func main() {
	// Basic input with Scan
	var pname string
	fmt.Println("Enter your favourite programming language: ")
	fmt.Scan(&pname)
	fmt.Println("So, your favourite programming language is", pname)

	// Multiple inputs with Scan
	var (
		name   string
		age    int
		gender rune
	)
	fmt.Println("Enter your name age and gender: ")
	fmt.Scan(&name, &age, &gender)
	fmt.Printf("Hello %s, you are a %c and %d years old", name, gender, age)
}
