package main

import (
	"flag"
	"fmt"
)

func main() {
	var port int
	var help bool
	flag.IntVar(&port, "p", 8000, "Provide a port number")
	flag.BoolVar(&help, "help", false, "Help")
	flag.Parse()
	if help {
		flag.PrintDefaults()
	} else {
		fmt.Println(port)
		vals := flag.Args()
		fmt.Println(vals)
	}
}
