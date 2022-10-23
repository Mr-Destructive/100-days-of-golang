package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func main() {

	f, err := os.Open("sample.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	reader := bufio.NewReader(f)
	chunk_size := 16
	chunk_list := []string{}
	buf := make([]byte, chunk_size)

	for {
		n, err := reader.Read(buf)

		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		chunk_list = append(chunk_list, string(buf[0:n]))
	}
	for _, chunk := range chunk_list {
		log.Print(chunk)
	}
}
