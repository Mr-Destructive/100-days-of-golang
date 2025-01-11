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
	ID     int      `json:"id"`
	Title  string   `json:"title"`
	Body   string   `json:"body"`
	UserId int      `json:"userId"`
}

func main() {
	URL := "https://jsonplaceholder.typicode.com/posts/4"
	req1, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(req1.Status)
	if req1.StatusCode != http.StatusOK {
		log.Fatal("Status code is not 200")
	}
	defer req1.Body.Close()
	body, err := io.ReadAll(req1.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Existing body:", string(body))

	reqBody := `{"body": "new body"}`
	fmt.Println("New body:", reqBody)

	req, err := http.NewRequest("PATCH", URL, strings.NewReader(reqBody))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

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

	err = json.Unmarshal(respBody, &updatedPost)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(updatedPost)
	fmt.Println("title:", updatedPost.Title)
	fmt.Println("body:", updatedPost.Body)
	fmt.Println("id:", updatedPost.ID)
	fmt.Println("user id:", updatedPost.UserId)
}
