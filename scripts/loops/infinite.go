package main

import "fmt"

func main() {
	//Inifinite Loop better keep commented
	/*
		flag := 4
		for {
			flag++
			fmt.Println(flag)
		}
	*/

	// Break Statement
	flag := 1
	for {
		fmt.Println(flag)
		flag++
		if flag == 7 {
			fmt.Println("It's time to break at", flag)
			break
		}
	}

	// Continue Statement
	counter := 2
	for counter < 4 {
		counter++
		if counter < 4 {
			continue
		}
		fmt.Println("Missed the Continue? at counter =", counter)
	}

}
