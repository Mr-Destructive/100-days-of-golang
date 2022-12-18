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
	f, err := os.OpenFile("test.txt", os.O_RDWR, 0660)
	/* test.txt
	   Hi
	   Hello
	   World
	   Gophers
	*/
	HandleError(err)
	m := bufio.NewScanner(f)
	bytes_till := 0
	//lines := []string{}
	// line to be appended
	lines_till := 2
	var lines_after string
	var line_till string
	i := 0
	for m.Scan() {
		line := m.Text()
		//lines = append(lines, line)
		if i < lines_till {
			bytes_till += bytes.Count([]byte(line), []byte{})
			if i > 0 {
				line_till += "\n"
			}
			line_till += line
		} else {
			lines_after += "\n" + line
		}
		i += 1
	}
	HandleError(m.Err())
	insert_text := line_till + "\nInserted content"
	insert_text_bytes := bytes.Count([]byte(insert_text), []byte{})
	lines_after_bytes := bytes.Count([]byte(lines_after), []byte{})

	err = os.WriteFile("test.txt", []byte(insert_text), 0660)
	HandleError(err)
	_, err = f.WriteAt([]byte(lines_after), int64(lines_after_bytes)+int64(insert_text_bytes))
	HandleError(err)
	defer f.Close()
}
