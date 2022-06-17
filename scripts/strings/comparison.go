package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "gopher"
	s2 := "Gopher"
	s3 := "gophy"

	isEqual := s1 == s2
	fmt.Printf("S1 and S2 are Equal? %t \n", isEqual)
	fmt.Println(s1 == s2)
	fmt.Println(s1 != s2)
	fmt.Println(s2 < s3)
	fmt.Println(s2 <= s3)
	fmt.Println(s2 > s3)
	fmt.Println(s2 >= s3)

	fmt.Println(strings.Compare(s1, s2))
	fmt.Println(strings.Compare(s1, s3))
	fmt.Println(strings.Compare(s2, s3))

	fmt.Println(strings.EqualFold(s1, s2))
	fmt.Println(strings.EqualFold(s1, s3))
	fmt.Println(strings.EqualFold(s2, s3))

}
