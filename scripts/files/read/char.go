package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	f, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanRunes)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	charlist := []byte{}
	for scanner.Scan() {
		char := byte(scanner.Text()[0])
		charlist = append(charlist, char)
	}
	log.Println(charlist)

}
