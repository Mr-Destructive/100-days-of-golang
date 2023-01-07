package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	fmt.Printf("Type of Args = %T\n", args)
	//fmt.Println(args[0], args[1])

	//total_args := len(os.Args[1:])
	//fmt.Println(total_args)

	for n, args := range os.Args {
		fmt.Println("Arg:", n, "->", args)
	}
	/*
			   // for excluding the 0th argument
		       for n, args := range os.Args[1:] {
			       fmt.Println("Arg:", n, "->", args)
			   }
	*/
	var port int
	var err error
	if len(os.Args) > 1 {
		port, err = strconv.Atoi(args[1])
		if err != nil {
			panic(err)
		}
	} else {
		port = 8000
	}
	fmt.Println(port)
}
