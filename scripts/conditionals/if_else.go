package main

import "fmt"

func main() {

	// basic if-else
	age := 16
	if age < 13 {
		fmt.Println("Kid")
	} else {
		fmt.Println("Teenager")
	}

	// else-if
	year := 2
	if year < 1 {
		fmt.Println("Freshman")
	} else if year == 2 {
		fmt.Println("Sophomore")
	} else if year == 3 {
		fmt.Println("Junior")
	} else if year == 4 {
		fmt.Println("Senior")
	} else {
		fmt.Println("Probably a Graduate")
	}
}
