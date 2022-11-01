package main

import (
	"log"
	"os"

	toml "github.com/pelletier/go-toml"
)

func main() {

	f, err := os.ReadFile("blog.toml")

	if err != nil {
		log.Fatal(err)
	}

	var data map[interface{}]interface{}

	err = toml.Unmarshal(f, &data)
	log.Println(data)

	if err != nil {
		log.Fatal(err)
	}

	for k, v := range data {
		log.Println(k, ":", v)
		switch t := v.(type) {
		case map[string]interface{}:
			for a, b := range t {
				log.Println(a, ":", b)
			}
		}
	}
}
