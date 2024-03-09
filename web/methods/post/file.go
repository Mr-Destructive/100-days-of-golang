package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

type ResponseFile struct {
	Files map[string]string `json:"files"`
}

func main() {
	// Define the URL of the API endpoint
	apiURL := "http://postman-echo.com/post"
	// file to uplaod
	fileName := "sample.csv"

	// Open the CSV file
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a new multipart writer
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add the file to the multipart form
	part, err := writer.CreateFormFile("csvFile", fileName)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		panic(err)
	}

	// Close the multipart writer
	writer.Close()

	// Make a POST request with the multipart form
	resp, err := http.Post(apiURL, writer.FormDataContentType(), body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Print response status code
	fmt.Println("Status Code:", resp.StatusCode)

	// Print response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	token := ResponseFile{}
	json.Unmarshal(respBody, &token)
	fmt.Println(token)
	fmt.Println(token.Files[fileName])
}
