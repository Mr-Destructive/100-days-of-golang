package main

import (
	"fmt"
	"net/http"
)

func main() {
	reqURL := "https://www.google.com"
	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
