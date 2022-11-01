package main

import (
	"log"
	"os"

	yaml "gopkg.in/yaml.v3"
)

func main() {

	f, err := os.ReadFile("blog.yaml")

	if err != nil {
		log.Fatal(err)
	}

	var data map[interface{}]interface{}

	err = yaml.Unmarshal(f, &data)
	log.Println(data)

	if err != nil {

		log.Fatal(err)
	}

	for k, v := range data {
		log.Println(k, ":", v)
		log.Printf("%T", v)
	}
}
