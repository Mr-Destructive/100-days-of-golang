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
	line_list := []string{}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		line_list = append(line_list, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	for _, line := range line_list {
		log.Println(line)
	}
}
