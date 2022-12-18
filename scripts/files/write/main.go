package main

import (
	"log"
	"os"
)

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Write
	str := "Hi, I am a gopher!\n"
	f, err := os.OpenFile("abc.txt", os.O_WRONLY, 0660)
	//f, err := os.Create("abc.txt")
	HandleError(err)
	_, err = f.Write([]byte(str))
	HandleError(err)
	langs := []string{"golang", "python", "rust", "javascript", "ruby"}
	for _, lang := range langs {
		_, err := f.WriteString(lang + "\n")
		HandleError(err)
	}
	defer f.Close()

	// Overwrite the pervious content in the file
	data := []byte{115, 111, 109, 101, 65}
	err = os.WriteFile("test.txt", data, 0660)
	HandleError(err)

	s := "Hello"
	err = os.WriteFile("test.txt", []byte(s), 0660)
	HandleError(err)

	// Append
	s = "\nHi\nWorld\nGophers"
	// open the file for writing/appending content
	f, err = os.OpenFile("test.txt", os.O_APPEND|os.O_WRONLY, 0660)
	HandleError(err)
	_, err = f.WriteString(s)
	HandleError(err)
	defer f.Close()

	// Append at a specific line
	// check out the file append.go

	//Delete conent from a file
	err = os.Truncate("test.txt", 5)
	HandleError(err)
}
