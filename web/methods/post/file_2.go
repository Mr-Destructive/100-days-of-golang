package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

type ResponseFile struct {
	Files map[string]string `json:"files"`
}

func main() {
	// Define the URL of the API endpoint
	apiURL := "http://postman-echo.com/post"
	// other dummy apis for uplaod
	// https://api.escuelajs.co/api/v1/files/upload
	// https://v2.convertapi.com/upload

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

	contentType := writer.FormDataContentType()
	fmt.Println(contentType)
	// Close the multipart writer
	writer.Close()
	fmt.Println(body.String())

	// Make a POST request with the multipart form
	resp, err := http.Post(apiURL, contentType, body)
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

	decodedFiles := make(map[string]string)
	for key, value := range token.Files {
		encodedVal := strings.Split(value, "base64,")[1]
		decodedContent, err := base64.StdEncoding.DecodeString(encodedVal)
		if err != nil {
			panic(err)
		}
		decodedFiles[key] = string(decodedContent)
	}
	fmt.Println(decodedFiles[fileName])
}
