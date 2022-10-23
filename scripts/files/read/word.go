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
	scanner.Split(bufio.ScanWords)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	wordlist := []string{}
	for scanner.Scan() {
		word := scanner.Text()
		wordlist = append(wordlist, word)
		log.Println(word)
	}
	log.Println(wordlist)

}
