package main

import (
	"fmt"
	"os"
)

func main() {
	file_name := "hello.txt"
	file, err := os.Open(file_name)
	if err != nil {
		fmt.Println("Error: ", err)
		//_, err := os.Create(file_name)
		//if err != nil {
		//	fmt.Println(err)
		//}
		//fmt.Println("Created File", file_name)
		//file, _ = os.Open(file_name)
	}
	fmt.Println(file.Stat())

}
