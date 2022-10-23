package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {

	f, err := os.Open("delimeter.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	reader := csv.NewReader(f)
	reader.Comma = ':'
	data, err := reader.ReadAll()
	if err != nil {

		log.Fatal(err)
	}
	log.Println(data)
}
