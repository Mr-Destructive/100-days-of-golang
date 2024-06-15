package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	apiURL := "https://reqres.in/api/users/5"

	req, err := http.NewRequest(http.MethodPut, apiURL, nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(respBody))

	defer resp.Body.Close()
}
