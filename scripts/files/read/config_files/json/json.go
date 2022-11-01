package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.ReadFile("blog.json")
	if err != nil {
		log.Println(err)
	}
	var result map[string]interface{}
	json.Unmarshal([]byte(f), &result)

	log.Println(result)
	for k, v := range result {
		fmt.Println(k, ":", v)
		fmt.Printf("%T\n", v)
	}

}
