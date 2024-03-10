package main

import (
	"fmt"
	"net/http"
	"bytes"
	"encoding/json"
)

func main() {
	apiURL := "https://reqres.in/api/users"
	bodyMap := map[string]interface{}{
		"name": "morpheus",
		"job":  "leader",
	}

	requestBody, err := json.Marshal(bodyMap)
	if err != nil {
		panic(err)
	}
	body := bytes.NewBuffer(requestBody)

	resp, err := http.Post(apiURL, "application/json", body)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	defer resp.Body.Close()
}
