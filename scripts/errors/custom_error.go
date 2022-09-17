package main

import (
	"fmt"
	"net/http"
)

type Invalid_URL_Error struct {
	message string
}

func (e Invalid_URL_Error) Error() string {
	return "Invalid URL : " + e.message
}

func get_resp(url_link string) (http.Response, error) {

	resp, err := http.Get(url_link)

	if err != nil {
		return http.Response{}, &Invalid_URL_Error{err.Error()}
	} else {
		defer resp.Body.Close()
		return *resp, nil
	}

}

func main() {

	url := "htts://meetgor.com"
	resp, err := get_resp(url)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}

}
