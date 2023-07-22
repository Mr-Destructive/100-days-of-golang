package main

import (
	"fmt"
)

func FindIndex[T comparable](arr []T, value T) int {
	for i, v := range arr {
		if v == value {
			return i
		}
	}
	return -1
}

type numeric interface {
	uint | uint8 | uint16 | uint32 | uint64 |
		int | int8 | int16 | int32 | int64 |
		float32 | float64
}

func Sum[T numeric](nums []T) T {
	var s T
	for _, n := range nums {
		s += n
	}
	return s
}

func main() {

	strSlice := []string{"m", "e", "e", "t"}
	fmt.Println(FindIndex(strSlice, "e"))
	fmt.Println(FindIndex(strSlice, "t"))
	fmt.Println(FindIndex(strSlice, "a"))

	intSlice := []int{10, 25, 33, 42, 50}
	fmt.Println(FindIndex(intSlice, 33))
	fmt.Println(FindIndex(intSlice, 90))

	intSlice = []int{10, 20, 30, 40, 50}
	fmt.Println(Sum(intSlice))
	floatSlice := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	fmt.Println(Sum(floatSlice))

}
