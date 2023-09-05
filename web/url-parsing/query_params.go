package main

import (
	"fmt"
	"net/url"
)

func main() {
	// a complex url with query params
	urlStr := "http://www.google.com/?q=hello+world&lang=en&q=gopher"
	parsedUrl, err := url.Parse(urlStr)
	if err != nil {
		panic(err)
	}
	fmt.Println(parsedUrl)
	fmt.Println(parsedUrl.Query())
	for k, v := range parsedUrl.Query() {
		fmt.Println("KEY:", k)
		for _, vv := range v {
			fmt.Println(vv)
		}
	}

	queryParams := parsedUrl.Query()

	// Get the value
	// will only give the first value
	fmt.Println(queryParams.Get("q"))

	// has value
	fmt.Println(queryParams.Has("q"))

	if queryParams.Has("lang") {
		fmt.Println(queryParams.Get("lang"))
	}

	//Add value to a key
	queryParams.Add("q", "ferris")
	fmt.Println(queryParams)

	// Encode the query params
	q := queryParams.Encode()
	fmt.Println(q)

	//Set the value
	queryParams.Set("q", "books")
	fmt.Println(queryParams)

	//Delete the value
	queryParams.Del("q")
	fmt.Println(queryParams)

}
