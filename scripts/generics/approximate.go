package main

import (
	"fmt"
)

type string2 string

type strings interface {
	~string
}

func PrintEach[T strings](arr T) {
	for _, v := range arr {
		fmt.Printf("%c\n", v)
	}
}

func main() {
	var s string2
	s = "hello"
	PrintEach(s)

}
