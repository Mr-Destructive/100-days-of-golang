package main

import "fmt"

func main() {
	//basic slice declaration
	var marks []int
	fmt.Println(marks)

	// declare and initialize using slice literal
	var frameworks = []string{"Django", "Laravel", "Flask", "Rails"}
	fmt.Println(frameworks)

	//using make function
	var langs = make([]string, 3, 5)
	langs[0], langs[1], langs[2] = "Python", "Go", "Javascript"
	fmt.Println(langs)
	fmt.Printf("Length = %d \nCapacity = %d\n", len(langs), cap(langs))
	langs = append(langs, "Java", "Kotlin", "PHP")
	fmt.Println(langs)
	fmt.Printf("Length = %d \nCapacity = %d\n", len(langs), cap(langs))

	// using new function
	langs2 := new([3]string)[0:2]

	langs2[0], langs2[1] = "Python", "Go"
	fmt.Println(langs2)
	fmt.Printf("Length = %d \nCapacity = %d\n", len(langs2), cap(langs2))
	langs2 = append(langs2, "Java", "Kotlin", "PHP")
	fmt.Println(langs2)
	fmt.Printf("Length = %d \nCapacity = %d\n", len(langs2), cap(langs2))

	// append function to add elements in slice
	var percentages = []float64{78.8, 85.7, 94.4, 79.8}
	fmt.Println(percentages)
	percentages = append(percentages, 60.5, 75.6)
	fmt.Println(percentages)

	// delete an elements
	marklist := []int{80, 85, 90, 75, 60}
	fmt.Println(marklist)
	var index int
	fmt.Printf("Enter the index to be deleted: ")
	fmt.Scan(&index)
	elem := marklist[index]
	marklist = append(marklist[:index], marklist[index+1:]...)
	fmt.Printf("The element %d was deleted.\n", elem)
	fmt.Println(marklist)

	// Using index slicing
	scores := []int{80, 85, 90, 75, 60, 56, 83}
	fmt.Println(scores)
	fmt.Println("From index 2 to 4", scores[2:5])
	fmt.Println("From index 0 to 2", scores[:3])
	fmt.Println("From index 3 to 6", scores[3:])

	// Modifying elements in slices
	word := []byte{'f', 'u', 'z', 'z', 'y'}
	fmt.Printf("%s\n", word)
	word[0] = 'b'
	word[len(word)-1] = 'z'
	fmt.Printf("%s\n", word)

	// three statement for loop in slices
	code := [7]rune{'g', 'o', 'l', 'a', 'n', 'g'}
	for i := 0; i < len(code); i++ {
		fmt.Printf("%c\n", code[i])
	}

	// using range based for loop
	for _, s := range scores[1:4] {
		fmt.Println(s)
	}

	start, i, end := 2, 2, 5
	modes := []string{"normal", "command", "insert", "visual", "select", "replace"}
	for range scores[start:end] {
		fmt.Printf("Element at index %d = %s \n", i, modes[i])
		i++
	}
}
