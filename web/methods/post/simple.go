package main

import (
	"fmt"
	"net/http"
)

func main() {
	apiURL := "https://reqres.in/api/users"
	// dummy api
	// few other apis
	// https://dummyjson.com/posts
	// https://jsonplaceholder.typicode.com/posts
	// https://httpbin.org/post
	// https://postman-echo.com/post

	// POST request
	resp, err := http.Post(apiURL, "application/json", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(resp)
	defer resp.Body.Close()
}
