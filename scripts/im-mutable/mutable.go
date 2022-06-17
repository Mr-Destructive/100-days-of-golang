package main

import "fmt"

func main() {

	// Data types which are Mutable
	// Mutable Type allow us to modify the underlying value of a memory address pointed by a variable
	// Slice
	s := []int{1, 2, 3}
	fmt.Printf("S[1] -> %p\n", &s[1])
	p := s
	p[1] = 4
	fmt.Printf("S[1] -> %p\n", &s[1])

	fmt.Println("s =", s)
	fmt.Println("p =", p)

	// Array
	a := [3]int{10, 20, 30}
	fmt.Printf("A[1] -> %p\n", &a[1])
	b := a
	b[1] = 40
	fmt.Printf("A[1] -> %p\n", &a[1])

	fmt.Println("a =", a)
	fmt.Println("b =", b)

	// Map
	m := map[string]int{"level": 5, "health": 9}
	fmt.Println(m)
	n := m
	n["food"] = 12

	fmt.Println("m =", m)
	fmt.Println("n =", n)

	// Channel
	// TODO Learn

}
