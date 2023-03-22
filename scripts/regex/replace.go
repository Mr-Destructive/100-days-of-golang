package main

import (
	"log"
	"regexp"
)

func main() {

	str := "Gophy gophers go in the golang grounds"
	exp := regexp.MustCompile(`Go|go`)
	log.Println("Original: ", str)
	replaced_str := exp.ReplaceAllString(str, "to")
	log.Println("Replaced: ", replaced_str)

	exp2 := regexp.MustCompile(`(Go|go)|(phers)|(rounds)`)
	log.Println(exp2.ReplaceAllString(str, "hop"))
	log.Println(exp2.ReplaceAllString(str, "$1"))
	log.Println(exp2.ReplaceAllString(str, "$2"))
	log.Println(exp2.ReplaceAllString(str, "$3"))
	log.Println(exp2.ReplaceAllString(str, "$1$2"))
	str = "Gophy gophers go in the golang cophers grounds"
	log.Println(exp2.ReplaceAllString(str, "$1$3"))

	log.Println(exp2.ReplaceAllLiteralString(str, "$1"))

}
