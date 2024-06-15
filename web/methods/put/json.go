package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	Name   string `json:"name"`
	Salary int    `json:"salary"`
	Age    string `json:"age"`
	ID     int    `json:"id,omitempty"`
}

type UserResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	user := User{
		Name:   "Alice",
		Salary: 50000,
		Age:    "25",
	}
	apiURL := "https://dummy.restapiexample.com/api/v1/update/11"

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
	req, err := http.NewRequest(http.MethodPut, apiURL, body)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	if resp.StatusCode == 429 {
		fmt.Println("too many requests")
        return
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(respBody))
	defer resp.Body.Close()

	// unmarshalling process
	// converting JSON to Go specific data structure/types
	var userResponse UserResponse
	if err := json.Unmarshal(respBody, &userResponse); err != nil {
		panic(err)
	}
	fmt.Println(userResponse)

}
