package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type ResponseLogin struct {
	Token string `json:"token"`
}

func main() {
	// Define the URL of the API endpoint
	apiURL := "https://reqres.in/api/login"

	// Define form data
	formData := url.Values{}
	formData.Set("email", "eve.holt@reqres.in")
	formData.Set("password", "cityslicka")

	// Encode the form data
    fmt.Println(formData.Encode())
	reqBody := strings.NewReader(formData.Encode())
	fmt.Println(reqBody)

	// Make a POST request with form data
	resp, err := http.Post(apiURL, "application/x-www-form-urlencoded", reqBody)
	if err != nil {
		panic(err) // Handle error
	}
	defer resp.Body.Close()

	// Print response status code
	fmt.Println("Status Code:", resp.StatusCode)

	// Print response body
	respBody, _ := io.ReadAll(resp.Body)
	token := ResponseLogin{}
	json.Unmarshal(respBody, &token)
	fmt.Println(token)
}
