package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name   string `json:"name"`
	Salary int    `json:"salary"`
	Age    int    `json:"age"`
}

/*
POST Request when you have a Go struct
Need to convert the Golang struct object to JSON
Read the JSON into a buffer / in-memeory
Send the POST request with that buffer as body
*/

func main() {
	user := User{
		Name:   "Alice",
		Salary: 50000,
		Age:    25,
	}
	apiURL := "https://dummy.restapiexample.com/api/v1/create"

	// marshalling process
	// converting Go specific data structure/types to JSON
	bodyBytes, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bodyBytes))

	// reading json into a buffer/in-memory
	body := bytes.NewBuffer(bodyBytes)

	// post request
	resp, err := http.Post(apiURL, "application/json", body)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(resp)
	defer resp.Body.Close()
}
