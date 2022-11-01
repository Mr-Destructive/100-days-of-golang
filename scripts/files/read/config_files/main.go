package main

import (
	"encoding/csv"
	"encoding/json"
	toml "github.com/pelletier/go-toml"
	yaml "gopkg.in/yaml.v3"
	"log"
	"os"
)

func check_error(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func read_json(file string) {
	f, err := os.ReadFile(file)
	check_error(err)
	var result map[string]interface{}
	json.Unmarshal([]byte(f), &result)

	log.Println(result)
	for k, v := range result {
		log.Println(k, ":", v)
		log.Printf("%T\n", v)
	}
}

func read_csv(file string) {
	f, err := os.Open(file)
	check_error(err)

	reader := csv.NewReader(f)

	n, err := reader.ReadAll()
	check_error(err)
	for _, line := range n {
		for _, text := range line {
			log.Println(text)
		}
	}
}

func read_yaml(file string) {
	f, err := os.ReadFile(file)
	check_error(err)
	var data map[string]interface{}

	err = yaml.Unmarshal(f, &data)
	check_error(err)
	log.Println(data)
	for k, v := range data {
		log.Println(k, ":", v)
		log.Printf("%T", v)
	}
}

func read_toml(file string) {
	f, err := os.ReadFile(file)

	check_error(err)

	var data map[interface{}]interface{}

	err = toml.Unmarshal(f, &data)
	log.Println(data)

	check_error(err)

	for k, v := range data {
		log.Println(k, ":", v)
		switch t := v.(type) {
		case map[string]interface{}:
			for a, b := range t {
				log.Println(a, ":", b)
				log.Printf("%T\n", b)
			}
		}
	}
}

func main() {
	read_json("json/blog.json")
	read_csv("csv/blog.csv")
	read_toml("toml/blog.toml")
	read_yaml("yaml/blog.yaml")
}
