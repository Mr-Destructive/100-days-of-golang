package main

import "fmt"

func main() {
    var percent = 30.738
    var f = 4.545

    // complex is used to store real and imaginary data
    var comp1 = complex(f, percent)
    var comp2 = complex(percent, f)
    fmt.Println(comp1 - comp2)
}
