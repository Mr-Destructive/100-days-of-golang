package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
	apiURL := "https://dummy.restapiexample.com/api/v1/update/4334"
	data := `{
        "employee_name": "tim",
        "employee_salary": "50000",
        "employee_age": 25
    }`
	body := bytes.NewBuffer([]byte(data))

	req, err := http.NewRequest(http.MethodPut, apiURL, body)
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
