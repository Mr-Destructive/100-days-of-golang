package main

import "fmt"

func main() {
	// Scanf for formatted input
	var (
		name   string
		age    int
		gender rune
	)
	fmt.Println("Enter your name age and gender: ")
	fmt.Scanf("%s \n %d %c", &name, &age, &gender)
	// for single input with Scanf
	// fmt.Scanf("%s", &name)
	fmt.Printf("Hello %s, you are a %c and %d years old", name, gender, age)
}
