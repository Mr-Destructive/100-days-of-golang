package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {

	// Using + operator
	s1 := "Go"
	s2 := "Programming"
	s3 := "30"

	s := s1 + s2 + s3
	fmt.Println(s)

	// Using += operator

	p := "Story"
	p += "Book"
	fmt.Println(p)

	// Using join function

	q := []string{"meetgor.com", "tags", "golang", "string"}
	r := strings.Join(q, "/")
	fmt.Println(r)

	// Using Sprintf function

	name := "peter"
	domain := "telecom"
	service := "ceo"

	email := fmt.Sprintf("%s.%s@%s.com", service, name, domain)
	fmt.Println(email)

	// Using Builder function

	c := []string{"j", "a", "v", "a"}
	var builder strings.Builder
	for _, item := range c {
		builder.WriteString(item)
	}

	fmt.Println(builder.String())

	// Using bytes buffer method

	var k bytes.Buffer

	for i := 0; i < 4; i++ {
		k.WriteString("no")
	}
	fmt.Println(k.String())

}
