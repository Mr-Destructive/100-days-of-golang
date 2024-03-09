package main

import (
	"bytes"
	"fmt"
	"net/http"
)

/*
POST Request when you have JSON string ready to send
Load that string into a buffer / in-memeory
Send the POST request with that buffer as body
*/
func main() {
	// dummy api
	apiURL := "https://dummy.restapiexample.com/api/v1/create"

	// json data
	data := `{
        "name": "Alice",
        "job": "Teacher"
    }`
	body := bytes.NewBuffer([]byte(data))

	// POST request
	resp, err := http.Post(apiURL, "application/json", body)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(resp)
	defer resp.Body.Close()
}
