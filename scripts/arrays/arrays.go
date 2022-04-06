package main

import "fmt"

func main() {
	// Simple Array declaration
	var languages [4]string
	languages[0] = "Python"
	fmt.Println(languages)

	// Using Array literal for declaration
	lang_array := [4]string{"Python", "Go", "Javascript", "C++"}
	fmt.Println(lang_array)

	// Accessing Array elements
	marks := [6]int{85, 89, 75, 93, 98, 60}
	fmt.Println(marks[1])
	fmt.Println(marks[5])
	fmt.Println(marks[3])

	// Modifying array elements
	name := [5]byte{'f', 'u', 'z', 'z', 'y'}
	fmt.Printf("%s\n", name)
	name[0] = 'b'
	name[4] = 'z'
	fmt.Printf("%s\n", name)

	// Letting Compiler figure out the length
	cart := [...]string{"Bag", "Shirt", "Watch", "Book"}
	fmt.Printf("There are %d items in your cart\n", len(cart))

	// Using len function to find length of array
	code := [7]rune{'#', '5', 'g', 't', 'm', 'y', '6'}
	fmt.Println("The length of the array is :", len(code))

	// Three statement for loop iteration of array
	for i := 0; i < len(code); i++ {
		fmt.Printf("%c\n", code[i])
	}

	// range based for loop to iterate over array
	for _, s := range cart {
		fmt.Println(s)
	}
}
