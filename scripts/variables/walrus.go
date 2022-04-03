package main

import "fmt"

func main() {
    place := "school"
    fmt.Println(place)

    x, y, z := "foo", 32, true
    fmt.Println(x, y, z)
    fmt.Printf("%T %T %T", x, y, z)
}
