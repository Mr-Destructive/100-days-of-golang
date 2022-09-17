package main

import (
	"fmt"
	"net/http"
	"net/url"
)

type Invalid_URL_Error struct {
	m string
}

func (e Invalid_URL_Error) Error() string {
	return "Invalid URL : " + e.m
}

func get_resp(url_link string) (http.Response, error) {
	_, err := url.ParseRequestURI(url_link)
	if err != nil {
		return http.Response{}, &Invalid_URL_Error{err.Error()}
	}
	resp, err := http.Get(url_link)
	if err != nil {
		return http.Response{}, &Invalid_URL_Error{err.Error()}
	} else {
		defer resp.Body.Close()
		return *resp, nil
	}

}

func main() {
	url := "https://.com"
	resp, err := get_resp(url)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
	url = "htt://meetgor.com"
	response, er := http.Get(url)
	if err != nil {
		fmt.Println(Invalid_URL_Error{er.Error()})
	} else {
		fmt.Println(response)
		defer response.Body.Close()
	}
}
