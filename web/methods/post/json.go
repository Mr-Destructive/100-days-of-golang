package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

/*
POST Request when you have Go struct
Need to convert the Golang struct object to JSON -> Marshal
Marshal is converting Go specific data structure/types to JSON
Load that string into a buffer / in-memeory
Send the POST request with that buffer as body

POST Request Response is in JSON format
But it is a in-memory object
Read the buffer into bytes
Convert bytes into JSON -> Unmarshal
Unmarshal is converting JSON into Go specific data structure
*/

type User struct {
	Name   string `json:"name"`
	Salary int    `json:"salary"`
	Age    string `json:"age"`
	ID     int    `json:"id,omitempty"`
}

type UserResponse struct {
	Status string `json:"status"`
	Data   User   `json:"data"`
}

func main() {
	user := User{
		Name:   "Alice",
		Salary: 50000,
		Age:    "25",
	}
	apiURL := "https://dummy.restapiexample.com/api/v1/create"

	// marshing process
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

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// unmarshing process
	// converting JSON to Go specific data structure/types
	var userResponse UserResponse
	if err := json.Unmarshal(respBody, &userResponse); err != nil {

		panic(err)
	}
	fmt.Println(userResponse)
	fmt.Println(userResponse.Data)

}
