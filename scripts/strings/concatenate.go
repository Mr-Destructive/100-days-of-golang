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
	fmt.Println("builder = ", builder.String())
	b := []byte{'s', 'c', 'r', 'i', 'p', 't'}
	for _, item := range b {
		builder.WriteByte(item)
	}
	fmt.Println("builder = ", builder.String())
	builder.WriteRune('s')
	fmt.Println("builder = ", builder.String())
	fmt.Println("builder = ", builder)

	// Using bytes buffer method

	var buf bytes.Buffer

	for i := 0; i < 2; i++ {
		buf.WriteString("ja")
	}
	fmt.Println(buf.String())

	buf.WriteByte('r')

	fmt.Println(buf.String())

	k := []rune{'s', 'c', 'r', 'i', 'p', 't'}
	for _, item := range k {
		buf.WriteRune(item)
	}
	fmt.Println(buf.String())
	fmt.Println(buf)

	buff := make([]byte, 0, 0)
	buff.WriteByte('s')
	buff.WriteByte('p')
	buff.off = 1
	fmt.Println(buff)

}
