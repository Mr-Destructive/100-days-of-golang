package main

import (
	"fmt"
	"regexp"
)

func main() {

	exp := regexp.MustCompile(`[a-zA-Z]*(\d*)[a-zA-Z]*@[a-zA-Z]*(\d*)[a-zA-Z]+\.[a-z|A-Z]{2,}`)
	str := "abc@def.com ab1e@def.com abe2@def.com a21be@def.com abe@de2f.com"
	matches := exp.FindAllStringSubmatch(str, -1)
	fmt.Println(matches)

}
