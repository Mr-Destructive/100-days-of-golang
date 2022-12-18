package main

import (
	"log"
	"os"
)

func main() {
	/*
	   Hi
	   Hello
	   World
	   Gophers
	*/
	// set only the first 6 bytes of the file
	// delete the rest
	err := os.Truncate("test.txt", 6)
	if err != nil {
		log.Fatal(err)
	}
	// Delete all the contents
	//err = os.Truncate("test.txt", 0)
	if err != nil {
		log.Fatal(err)
	}
}
