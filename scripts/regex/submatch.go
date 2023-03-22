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

	str := "abe21@def.com is the mail address of 8th user with id 124"
	//exp := regexp.MustCompile(`([a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z]{2,}:\d+)|(email|mail)|(\d{1,3})`)
	exp := regexp.MustCompile(
		`([a-zA-Z]+(\d*)[a-zA-Z]*@[a-zA-Z]+(\d*)[a-zA-Z]*\.[a-z|A-Z]{2,})` +
			`|(mail|email)` +
			`|(\d{1,3})`,
	)

	match := exp.FindStringSubmatch(str)
	match_index := exp.FindStringSubmatchIndex(str)
	match_indexes := exp.FindAllStringSubmatchIndex(str, -1)
	log.Println(match)
	log.Println(match_index)
	log.Println(match_indexes)
	matches := exp.FindAllStringSubmatch(str, -1)
	log.Println(matches)

	email_file, err := os.ReadFile("subtext.txt")
	log_error(err)
	mail_match := exp.FindSubmatch(email_file)
	log.Printf("%T\n", mail_match)
	log.Printf("%s\n", mail_match)
	mail_matches := exp.FindAllSubmatch(email_file, -1)
	//log.Println(mail_matches)
	log.Printf("%T\n", mail_matches)
	log.Printf("%s\n", mail_matches)
}
