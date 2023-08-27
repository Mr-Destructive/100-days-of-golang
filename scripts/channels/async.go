package main

import (
	"fmt"
	"net/http"
)

func fetchURL(url string, ch chan<- http.Response) {
	go func() {
		res, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()
		ch <- *res
	}()
}

func task(name string) {
	fmt.Println("Task", name)
}

func main() {
	fmt.Println("Start")

	url := "http://google.com"

	respCh := make(chan http.Response)

	fetchURL(url, respCh)

	task("A")
	task("B")

	response := <-respCh
	fmt.Println("Response Status:", response.Status)

	fmt.Println("Done")
}
