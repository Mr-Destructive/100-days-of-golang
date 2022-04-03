package main

import (
    "fmt"
)

func main() {
    // byte is alias for uint8
    var s byte = 't'
    fmt.Println(s)

    // %c is the placeholder for character 
    const c byte = 't'
    fmt.Printf("Character = %c \nInteger value = %d\n", c, c)
}
