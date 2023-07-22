package main

import (
	"fmt"
)

func ForEach[T any](arr []T, f func(T)) {
	for _, v := range arr {
		f(v)
	}
}

func main() {

	strSlice := []string{"b", "e", "a"}
	ForEach(strSlice, func(v string) {
		fmt.Println(v)
	})

	slice := []int{10, 25, 33, 42, 50}
	var evenSlice []int
	ForEach(slice, func(v int) {
		isEven := v%2 == 0
		if isEven {
			evenSlice = append(evenSlice, v)
		}
	})
	fmt.Println(evenSlice)

}
