package main

import (
	"flag"
	"fmt"
)

func main() {
	port := flag.Int("p", 8000, "Provide a port number")
	dir := flag.String("dir", "output_dir", "Directory")
	publish := flag.Bool("publish", false, "Publish the article")
	help := flag.Bool("help", false, "Help")
	if *help {
		flag.PrintDefaults()
	} else {
		fmt.Println(*port)
		fmt.Println(*dir)
		flag.Set("dir", "dumps")
		fmt.Println(*dir)
		fmt.Println(flag.NFlag())
		fmt.Println(flag.NArg())
		fmt.Println(*publish)
		if *publish {
			fmt.Println("Publishing article...")
		} else {
			fmt.Println("Article saved as Draft!")
		}
		vals := flag.Args()
		fmt.Println(vals)
	}
}
