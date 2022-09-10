package main

import (
	"fmt"
)

func compute_length(n interface{}) int {
	switch n.(type) {
	case int:
		return (n).(int)
	case string:
		return len(n.(string))
	case []int:
		return len(n.([]int))
	case []string:
		return len(n.([]string))
	case []float64:
		return len(n.([]float64))
	default:
		return n.(int)
	}
}

func main() {
	num := compute_length(2)
	fmt.Println(num)
	num = compute_length("hello world")
	fmt.Println(num)
	num = compute_length([]int{1, 3, 4})
	fmt.Println(num)
	num = compute_length([]string{"python", "c++", "js", "java", "go"})
	fmt.Println(num)
}
