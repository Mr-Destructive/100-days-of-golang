package main

import (
    "fmt"
)

func main() {
    var code = "12AB34CD"
    fmt.Println(code[6])
    fmt.Printf("2nd Character in string = %c\n\n", code[4])

    var statement = `This is the first line
The next line
The last line`

    fmt.Println(statement)
}
