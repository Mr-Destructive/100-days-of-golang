package main

import "fmt"

func swap(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

type Book struct {
	pages int
	genre string
	title string
}

func main() {

	// pointer declaration
	var ptr *int
	fmt.Println(ptr)

	// accessing memory address of variable with &
	n := 34
	var a_pointer *int = &n
	fmt.Println(a_pointer)

	// de-referncing pointers with *
	m := *a_pointer
	fmt.Println(m)

	//pass by reference to a function
	x := 3
	y := 6
	k := &x
	p := &y
	fmt.Printf("Before swapping : x = %d and y = %d.\n", x, y)
	swap(k, p)
	fmt.Printf("After swapping  : x = %d and y = %d.\n", x, y)

	// pointer to a struct instance
	new_book := Book{120, "fiction", "Harry Potter"}
	fmt.Println(new_book)
	fmt.Printf("Type of new_book -> %T\n", new_book)
	book_ptr := &new_book
	book_ptr.title = "Games of Thrones"
	fmt.Println(new_book)
}
