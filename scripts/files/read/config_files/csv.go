package main

import (
	"encoding/csv"
	"log"
	"net/http"
)

func main() {

	url := "https://github.com/woocommerce/woocommerce/raw/master/sample-data/sample_products.csv"
	response, err := http.Get(url)

	if err != nil {
		log.Println(err)
		return
	}

	defer response.Body.Close()
	reader := csv.NewReader(response.Body)
	n, err := reader.ReadAll()
	if err != nil {
		log.Println(err)
		return
	}
	for _, line := range n {
		for _, text := range line {
			log.Println(text)
		}
	}
}
