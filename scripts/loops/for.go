package main

import "fmt"

func main() {

	//three statement basic for loop
	for k := 0; k < 4; k++ {
		fmt.Println(k)
	}

	// range based for loop
	name := "GOLANG"
	for i, s := range name {
		fmt.Printf("%d -> %c\n", i, s)
	}
}
