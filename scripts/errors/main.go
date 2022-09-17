package main

import (
	"fmt"
)

func main() {
	var s string
	n, err := fmt.Scanf("%s", &s)
	if err != nil {
		fmt.Println(err)
		//panic(err)
	} else {
		fmt.Println(n)
		if s[0] == 'a' {
			fmt.Println(s)
		}
	}
}
