package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	baseURL := "https://reqres.in"
	userID := 2
	apiURL := fmt.Sprintf("%s/api/users/%d", baseURL, userID)

	req, err := http.NewRequest(http.MethodDelete, apiURL, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Response Status:", resp.Status)
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Response Body:", string(respBody))
}
