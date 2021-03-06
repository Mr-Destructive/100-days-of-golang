package main

import "fmt"

func main() {

	// Data types which are Immutable
	// String
	str := "gopher"
	str_copy := str
	str_copy = "cooper"
	fmt.Println("str = ", str)
	fmt.Println("str_copy = ", str_copy)

	// Character at index cannot be changed in string

	s := "StarWars"
	s[4] = 'C'
	// s[4] = "C"
	// also won't work
	fmt.Println(s)

	// bool
	boolean := true
	b := boolean
	b = false
	fmt.Println("boolean = ", boolean)
	fmt.Println("b = ", b)

	// pointer, function pointers
	n := 567
	t := 123
	ptr := &n
	ptr_new := ptr
	fmt.Println("ptr = ", ptr)
	fmt.Println("ptr_new = ", ptr_new)

	ptr_new = &t

	fmt.Println("ptr = ", ptr)
	fmt.Println("ptr_new = ", ptr_new)

	// Interface
	// TODO Learn

}
