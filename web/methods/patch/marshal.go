package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type Post struct {
	ID     int    `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Body   string `json:"body,omitempty"`
	UserId int    `json:"userId,omitempty"`
}

func main() {

	// define URL to hit the API
	baseURL := "https://jsonplaceholder.typicode.com"
	postId := 4
	postURL := fmt.Sprintf("%s/posts/%d", baseURL, postId)

	userObj := Post{
		Body: "New Body",
	}

	var reqBody []byte
	reqBody, err := json.Marshal(userObj)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("New body:", string(reqBody))

	// send a new request, with the `PATCH` method, url and the body
	req, err := http.NewRequest("PATCH", postURL, strings.NewReader(string(reqBody)))
	if err != nil {
		log.Fatal(err)
	}
	// set the header content type to json
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println("Response status code:", resp.StatusCode)
	fmt.Println("Response Status:", resp.Status)

	var updatedPost Post

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// convert the response json bytes to Post object in golang
	err = json.Unmarshal(respBody, &updatedPost)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(updatedPost)
	fmt.Println(updatedPost.Title)

}
