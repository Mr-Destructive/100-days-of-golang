package main

import (
	"fmt"
	"strings"
)

func main() {

	s1 := "gopher"
	s2 := "Gopher"
	s3 := "gopher"

	isEqual := s1 == s2

	//"gopher" is not same as "Gopher" and hence `false`
	fmt.Printf("S1 and S2 are Equal? %t \n", isEqual)
	fmt.Println(s1 == s2)

	// "gopher" is not equal to "Gopher" and hence `true`
	fmt.Println(s1 != s2)

	// "Gopher" comes first lexicographically than "gopher" so return true
	// G -> 71 in ASCII and g -> 103
	fmt.Println(s2 < s3)
	fmt.Println(s2 <= s3)

	// "Gopher" is not greater than "gopher" as `G` comes first in ASCII table
	// So G value is less than g i.e. 71 > 103 which is false
	fmt.Println(s2 > s3)
	fmt.Println(s2 >= s3)

	fmt.Println(strings.Compare(s1, s2))
	fmt.Println(strings.Compare(s1, s3))
	fmt.Println(strings.Compare(s2, s3))

	fmt.Println(strings.EqualFold(s1, s2))
	fmt.Println(strings.EqualFold(s1, s3))
	fmt.Println(strings.EqualFold(s2, s3))

}
