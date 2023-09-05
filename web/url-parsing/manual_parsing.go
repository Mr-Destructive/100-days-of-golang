package main

import (
	"fmt"
	"strings"
)

func main() {
	urlString := []string{"http://www.google.com",
		"http://www.google.com/about/",
		"http://www.google.com/about?q=hello",
		"http://www.google.com:8080/about?q=hello",
		"http://user:password@example.com:8080/path/to/resource?query=value#fragment",
	}
	for _, url := range urlString {
		hostStr := strings.Split(url, "//")[1]
        hostName := strings.Split(hostStr, "/")[0]
		fmt.Println(hostName)
	}
}
