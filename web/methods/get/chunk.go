package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	reqURL := "https://httpbin.org/html"
	resp, err := http.Get(reqURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

    // create a empty buffer
	buf := new(bytes.Buffer)

    // create a chunk buffer of a fixed size
	chunk := make([]byte, 1024)

	for {
		// Read into buffer
		n, err := resp.Body.Read(chunk)
		if err != nil {
			break
		}
        // append the chunk to the buffer
		buf.Write(chunk[:n])
		fmt.Printf("%s\n", chunk[:n])
	}

    // entire body stored in bytes
	fmt.Println(buf.String())
}
