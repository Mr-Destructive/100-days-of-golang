package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
	apiURL := "https://dummy.restapiexample.com/api/v1/create"

	data := `{
        "name": "jim",
        "salary": "50000",
        "age": 25,
    }`
	body := bytes.NewBuffer([]byte(data))

	// POST request
	resp, err := http.Post(apiURL, "application/json", body)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(resp)
	// read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(respBody))
	defer resp.Body.Close()
}
