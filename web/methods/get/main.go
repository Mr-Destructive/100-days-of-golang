package main

import (
	"fmt"
	"net/http"
)

func main() {
	reqURL := "https://www.google.com"
	resp, err := http.Get(reqURL)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
	fmt.Println(resp.Status)
	fmt.Println("Status Code:", resp.StatusCode)
}
