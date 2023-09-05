package main

import (
	"fmt"
	"net/url"
	"strings"
)

func main() {

	// url with multiple path
	url1 := "https://github.com/appwrite/console/issues/396"
	parsedUrl, err := url.Parse(url1)
	if err != nil {
		panic(err)
	}
	fmt.Println(parsedUrl)
	fmt.Println("Paths", parsedUrl.Path)
	paths := strings.Split(parsedUrl.Path, "/")
	for _, path := range paths {
		fmt.Println(path)
	}

	// url with subdomain
	url2 := "https://docs.mindsdb.com/what-is-mindsdb"
	parsedUrl, err = url.Parse(url2)
	if err != nil {
		panic(err)
	}
	fmt.Println(parsedUrl)
	fmt.Println("Host", parsedUrl.Host)
}
