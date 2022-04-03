package main

import "fmt"

func main() {

    var (
        x int8 = 100
        y byte = '#'
        z =  "daysofcode"
    )

    fmt.Printf(" x = %d \n y = %c \n z = %s \n",x,y,z)

    var pi, e, G float32 = 3.141, 2.718, 6.67e-11
    var start, end byte = 65, 90
    fmt.Println(pi, e, G)
    fmt.Printf("%c %c\n",start, end)
}
