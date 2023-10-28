package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	req, err := http.NewRequest(http.MethodGet, "https://api.github.com/users/mr-destructive", nil)
	if err != nil {
		panic(err)
	}
	// if the below line is commented, it works fine, but for real authentication-based APIs/services, it won't work
	//req.Header.Add("Authorization", "token YOUR_TOKEN")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
