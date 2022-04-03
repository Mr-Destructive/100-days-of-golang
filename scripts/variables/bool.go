package main

import "fmt"

func main() {
    // Default value of bool is false
    var power bool
    fmt.Println(power)

    // %t is the placeholder for bool
    const result = true
    fmt.Printf("The statement is %t", result)
}
