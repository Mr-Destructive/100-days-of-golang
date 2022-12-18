package main

import (
	"fmt"
	"log"
	"os"
)

func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	f, err := os.Create("temp.txt")
	HandleErr(err)
	name, lang, exp := "John", "go", 2
	_, err = fmt.Fprint(f, "Hi, I am ", name, "\n")
	HandleErr(err)
	_, err = fmt.Fprintf(f, "Language of choice: %s.\n", lang)
	HandleErr(err)
	_, err = fmt.Fprintln(f, "Having", exp, "years of experience.")
	HandleErr(err)
	defer f.Close()
}
