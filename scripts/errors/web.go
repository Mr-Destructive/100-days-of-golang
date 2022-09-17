package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := "htts://meetgor.com/"
	resp, err := http.Get(url)
	fmt.Println("URL:", url)
	if err != nil {
		fmt.Println(err)
		panic(err)
	} else {
		fmt.Println(resp.StatusCode)
	}
}

/*
	if resp, err := http.Get(url); err != nil {
		fmt.Println(err)
		panic(err)
	} else {
		fmt.Println(resp.StatusCode)
	}
*/
