package main

import (
	"flag"
	"fmt"
)

func main() {
	var port int
	var dir string
	var publish bool
	var help bool
	flag.IntVar(&port, "p", 8000, "Provide a port number")
	flag.StringVar(&dir, "dir", "output_dir", "Directory")
	flag.BoolVar(&publish, "publish", false, "Publish the article")
	flag.BoolVar(&help, "help", false, "Help")
	flag.Parse()
	if help {
		flag.PrintDefaults()
	} else {
		fmt.Println(port)
		fmt.Println(dir)
		flag.Set("dir", "dumps")
		fmt.Println(dir)
		fmt.Println(flag.NFlag())
		fmt.Println(flag.NArg())
		fmt.Println(publish)
		if publish {
			fmt.Println("Publishing article...")
		} else {
			fmt.Println("Article saved as Draft!")
		}
		vals := flag.Args()
		fmt.Println(vals)
	}
}
