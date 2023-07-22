package main

import (
	"fmt"
)

func Print[T any](stuff T) {
	fmt.Println(stuff)
}

type Response[T any] struct {
	StatusCode int
	Data       T
}

func main() {

	Print("hello")
	Print(123)
	Print(3.148)

	resp := Response[string]{StatusCode: 200, Data: "hello"}
	Print(resp)
	resp2 := Response[float64]{StatusCode: 200, Data: 353.23376}
	Print(resp2)

}
