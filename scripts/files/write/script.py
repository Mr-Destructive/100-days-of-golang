package main

import (
	"bufio"
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
	// Write
	// Overwrite the pervious content in the file
	//data := []byte{115, 111, 109, 101, 65}
	//err := os.WriteFile("test.txt", data, 0660)
	//HandleError(err)

	//s := "Hello"
	//err = os.WriteFile("test.txt", []byte(s), 0660)
	//HandleError(err)

	//// Append
	//s = "\nWorld\nGophers"
	//// open the file for writing/appending content
	//f, err := os.OpenFile("test.txt", os.O_APPEND|os.O_WRONLY, 0660)
	//HandleError(err)
	//_, err = f.WriteString(s)
	//HandleError(err)
	//defer f.Close()

	f, err := os.OpenFile("test.txt", os.O_RDWR, 0660)
	HandleError(err)
	m := bufio.NewScanner(f)
	bytes_till := 0
	lines := []string{}
	lines_till := 1
	lines_after := ""
	i := 0
	for m.Scan() {
		line := m.Text()
		lines = append(lines, line)
		if i < lines_till {
			bytes_till += bytes.Count([]byte(line), []byte{})
		} else {
			lines_after += "\n" + line
		}
		i += 1
	}
	log.Println(bytes_till)
	//lines_before := lines[:lines_till]
	//lines_after := lines[lines_till:]
	//log.Println(lines_before)
	log.Println(lines_after)
	HandleError(m.Err())
	insert_text := "\nInserted content"
	insert_text_bytes := bytes.Count([]byte(insert_text), []byte{})
	lines_after_bytes := bytes.Count([]byte(lines_after), []byte{})

	_, err = f.WriteAt([]byte(insert_text), int64(bytes_till)+1)
	HandleError(err)
	_, err = f.WriteAt([]byte(lines_after), int64(lines_after_bytes)+int64(insert_text_bytes))
	HandleError(err)
	defer f.Close()
}

/*
Hi
Hello
World
Gophers
*/
