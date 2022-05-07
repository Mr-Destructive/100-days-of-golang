package main

import "fmt"

func main() {
	a := 45
	b := "Something"
	fmt.Println(a > 40 && b == "Something")
	fmt.Println(a < 40 && b == "Something")
	fmt.Println(a < 40 || b == "Something")
	fmt.Println(a < 40 || b != "Something")
	fmt.Println(!(a < 40 || b != "Something"))
}
