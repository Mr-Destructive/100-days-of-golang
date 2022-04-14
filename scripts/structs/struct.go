package main

import "fmt"

type Article struct {
	title        string
	is_published bool
	words        int
}

func main() {
	golang := Article{"Golang Intro", true, 2000}
	fmt.Println(golang)

	vim := Article{title: "Vim: Keymapping", is_published: false}
	fmt.Println(vim)

	django := *new(Article)
	fmt.Println(django)

	django.title = "Django View and URLs"
	django.words = 3500
	django.is_published = true
	fmt.Println(django)
}
