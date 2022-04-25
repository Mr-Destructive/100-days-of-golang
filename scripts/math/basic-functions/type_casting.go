package main

import "fmt"

func main() {
	foo := 77
	fmt.Printf("Type of foo = %T \n", foo)
	fmt.Println("foo = ", int(foo))
	fmt.Println("String Cast: ", string(foo))
	fmt.Println("Float Cast: ", float64(foo))
	var a float64 = 5
	var b float64 = 0
	fmt.Println(a / b)
}
