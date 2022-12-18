package main

import (
	"bytes"
	"log"
	"os"
)

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	filename := "test.txt"
	file, err := os.ReadFile(filename)
	HandleError(err)
	old_text := "e"
	new_text := "E"
	new_content := bytes.Replace(file, []byte(old_text), []byte(new_text), 2)
	err = os.WriteFile(filename, new_content, 0660)
	HandleError(err)
}
