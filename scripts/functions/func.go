package main

import "fmt"

func main() {

	// basic hello world function
	hello_world()

	// pass parameter
	greet_me("Meet")
	n := "John"
	greet_me(n)

	// return value
	y := line_eq(3, 1, 2)
	fmt.Println("for x = 3 , y = ", y)

	// multiple return values
	s, c, odd := sqube(5)
	fmt.Println("for x = 5 , x^2 =", s, "x^3 =", c)
	if odd == true {
		fmt.Println("x is odd")
	} else {
		fmt.Println("x is true")
	}
}

func hello_world() {
	name := "Gopher"
	fmt.Println("Hello", name)
}

func greet_me(name string) {
	fmt.Println("Hello,", name, "!")
}

func line_eq(x int, m int, c int) int {
	return ((m * x) + c)
}

func sqube(x int) (int, int, bool) {
	square := x * x
	cube := x * x * x
	var is_odd bool
	if x%2 == 0 {
		is_odd = false
	} else {
		is_odd = true
	}
	return square, cube, is_odd
}
