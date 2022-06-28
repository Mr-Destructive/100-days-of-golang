package main

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	s1 := "Ubuntu 22"
	s2 := "meet"
	s3 := "IND"

	// ToLower method for converting string into Lower case
	fmt.Println(strings.ToLower(s1))
	fmt.Println(strings.ToLower(s2))
	fmt.Println(strings.ToLower(s3))

	// ToUpper method for converting string into Upper case
	fmt.Printf("\n")
	fmt.Println(strings.ToUpper(s1))
	fmt.Println(strings.ToUpper(s2))
	fmt.Println(strings.ToUpper(s3))

	// Title method for converting string into title case
	fmt.Printf("\n")
	cases := cases.Title(language.English)
	fmt.Println(cases.String(s1))
	fmt.Println(cases.String(s2))
	fmt.Println(cases.String(s3))

	str := "javascript"
	substr := "script"
	s := "python"

	// Contains method for finding substring in the string
	fmt.Println(strings.Contains(str, substr))
	fmt.Println(strings.Contains(str, s))

	// ContainsAny method for finding characters in the string
	fmt.Println(strings.ContainsAny(str, "joke"))
	fmt.Println(strings.ContainsAny(str, "xyz"))
	fmt.Println(strings.ContainsAny(str, ""))

	// Split method for splitting string into slice of string
	link := "meetgor.com/blog/golang/strings"
	fmt.Println(strings.Split(link, "/"))
	fmt.Println(strings.SplitAfter(link, "/"))

	// SplitAfter method for splitting string into slice of string with the pattern
	data := "200kms50kms120kms"
	fmt.Println(strings.Split(data, "kms"))
	fmt.Println(strings.SplitAfter(data, "kms"))

	// Repeat method for creating strings with given string and integer
	pattern := "OK"
	fmt.Println(strings.Repeat(pattern, 3))

	// Fields method for extracting string from the given string with white space as delimiters
	text := "Python is a prgramming language.   Go is not"
	text_data := strings.Fields(text)
	fmt.Println(text_data)
	for _, d := range text_data {
		fmt.Println("data = ", d)
	}
}
