package main

import "fmt"

func main() {

	//basic switch statement
	age := 11
	var result string
	switch {
	case age < 13:
		result = "Kid"
	case age < 20:
		result = "Teenager"
	case age < 25:
		result = "Adult"
	case age < 35:
		result = "Senior"
	}
	fmt.Printf("The person is a %s with age %d.\n", result, age)

	//default case
	language := "go"
	var devs string
	switch language {
	case "go":
		devs = "gopher"
	case "rust":
		devs = "rustacean"
	case "python":
		devs = "pythonista"
	case "java":
		devs = "Duke"
	default:
		language = "javascript"
		devs = "developer"
	}
	fmt.Println("A person who codes in", language, "is called a", devs)

	// fallthrough

	character := 'a'
	fmt.Printf("The input character is = %c\n", character)
	switch {
	case character == 97:
		fmt.Printf("Its %c\n", character)
		fallthrough
	case character < 107 && character > 97:
		fmt.Println("It's between a and k")
		fallthrough
	case character < 117 && character > 97:
		fmt.Println("It's between a and t")
		fallthrough
	case character < 122 && character > 97:
		fmt.Println("It's between a and u")
	default:
		fmt.Println("Its not a lowercase alphabet")
	}

}
