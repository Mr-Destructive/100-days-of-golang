package main

import (
	"encoding/xml"
	"log"
	"os"
)

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

func main() {

	//f, err := os.Open("rss.xml")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer f.Close()
	//log.Printf("%T\n", f)
	//data, err := ioutil.ReadAll(f)
	//if err != nil {
	//	log.Fatal(err)
	//}
	f, err := os.ReadFile("rss.xml")
	d := Channel{}
	err = xml.Unmarshal(f, &d)
	log.Println(d)
	if err != nil {
		log.Printf("error: %v", err)
	}
	for _, item := range d.Item {
		log.Println(item.Title)
	}
}
