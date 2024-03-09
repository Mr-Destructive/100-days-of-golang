package main

import (
	"fmt"
	"net/http"
)

func main() {
	apiURL := "https://reqres.in/api/users"

	req, err := http.NewRequest(http.MethodPost, apiURL, nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(resp)
	defer resp.Body.Close()
}
