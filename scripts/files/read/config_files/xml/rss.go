package main

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"
)

type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	XMLName     xml.Name `xml:"channel"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Item        []Item   `xml:"item"`
}

type Item struct {
	XMLName xml.Name `xml:"item"`
	Title   string   `xml:"title"`
	Link    string   `xml:"link"`
}

func check_error(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	url := "https://meetgor.com/rss.xml"
	response, err := http.Get(url)

	check_error(err)

	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)

	check_error(err)

	d := Rss{}
	err = xml.Unmarshal(data, &d)

	check_error(err)

	for _, item := range d.Channel.Item {
		log.Println(item.Title)
	}
}
