package main

import (
	"fmt"
)

type Stack[T any] struct {
	items []T
}

// it is not needed, but handy and a good practise to create one
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
	if len(s.items) == 0 {
		panic("Stack is empty")
	}
	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item
}

func main() {
	intStack := NewStack[int]()
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)
	fmt.Println("Integer Stack")
	fmt.Println(intStack)
	intStack.Pop()
	intStack.Pop()
	fmt.Println(intStack)

	// without the NewStack method
	strStack := Stack[string]{}
	strStack.Push("c")
	strStack.Push("python")
	strStack.Push("mojo")
	fmt.Println("String Stack:")
	fmt.Println(strStack)
	strStack.Pop()
	fmt.Println(strStack)
}
