package main

import (
	"log"
	"os"
)

func main() {

	text, err := os.ReadFile("sample.txt")

	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(text))
}
