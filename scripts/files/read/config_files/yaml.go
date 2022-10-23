package main

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v3"
)

func main() {

	f, err := ioutil.ReadFile("blog.yaml")

	if err != nil {
		log.Fatal(err)
	}

	data := make(map[interface{}]interface{})

	err = yaml.Unmarshal(f, &data)
	log.Println(data)

	if err != nil {

		log.Fatal(err)
	}

	for k, v := range data {
		log.Println(k, ":", v)
	}
}
