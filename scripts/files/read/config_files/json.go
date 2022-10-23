package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func main() {

	f, err := ioutil.ReadFile("blog.json")
	if err != nil {
		log.Println(err)
	}
	var result map[string]interface{}
	json.Unmarshal([]byte(f), &result)

	log.Println(result)
	log.Println(result["title"])

}
