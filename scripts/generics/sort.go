package main

import (
	"fmt"
)

type sortable interface {
	int | float64 | string
}

func sort[T sortable](arr []T) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func main() {
	arr := []int{10, 28, 57, 96, 76, 8, 97, 10}
	sort(arr)
	fmt.Println(arr)
	strarr := []string{"b", "e", "a"}
	sort(strarr)
	fmt.Println(strarr)
}
