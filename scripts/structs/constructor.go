package main

import "fmt"

type Repository struct {
	name       string
	file_count int
}

func New_Repository(name string, file_count int) *Repository {
	file_count++
	name = "Test"
	return &Repository{name, file_count}
}

func main() {
	blog := *New_Repository("", 0)
	fmt.Println(blog)
}
