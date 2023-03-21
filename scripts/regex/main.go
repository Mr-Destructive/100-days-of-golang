package main

import (
	"log"
	"os"
	"regexp"
)

func log_error(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	str := "gophers of the goland"
	is_matching, err := regexp.MatchString("go", str)
	log_error(err)
	log.Println(is_matching)
	file_content, err := os.ReadFile("temp.txt")
	log_error(err)
	is_matching, err = regexp.Match(`.*memory`, file_content)
	log_error(err)
	log.Println(is_matching)
	is_matching, err = regexp.Match(`text `, file_content)
	log_error(err)
	log.Println(is_matching)

	exp, err := regexp.Compile(`\b\d{5}(?:[-\s]\d{4})?\b`)
	log_error(err)
	pincode_file, err := os.ReadFile("pincode.txt")
	log_error(err)
	match := exp.FindAll(pincode_file, -1)
	matches := exp.FindAllString(string(pincode_file), -1)
	log.Println(match)
	log.Println(matches)

	match_indexes := exp.FindAllIndex(pincode_file, -1)
	match_index := exp.FindIndex(pincode_file)
	log.Println(match_indexes)
	log.Printf("%T\n", match_indexes)
	log.Println(match_index)
	if len(match_index) > 0 {
		log.Printf("%s\n", pincode_file[match_index[0]:match_index[1]])
	}
	log.Printf("%T\n", match_index)

}
