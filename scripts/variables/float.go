package main

import (
    "fmt"
)

func main() {

    // default type of decimal numbers is float64
    const percent = 30.5
    fmt.Println(percent+50)
    fmt.Printf("type = %T\n", percent)

    // optionally we can assign float32 for less precission than float64
    const percentage float32 = 46.34
    fmt.Println(percentage - 3.555)
    fmt.Printf("type = %T\n", percentage)
}
