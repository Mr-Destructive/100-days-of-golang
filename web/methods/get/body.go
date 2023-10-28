package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	reqURL := "https://httpbin.org/html"
	resp, err := http.Get(reqURL)
	if err != nil {
		panic(err)
	}
	// close the body object before returning of the function
	// this is to avoid the memory leak
	defer resp.Body.Close()

	// stream the data from the response body only once
	// it is not buffered in the memory
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
