package main

import "fmt"

func main() {
    var languages[4]string
    languages[0] = "Python"
    fmt.Println(languages)

    lang_array := [4]string {"Python", "Go", "Javascript", "C++"}
    fmt.Println(lang_array)

    marks := [6]int {85, 89, 75, 93, 98, 60}
    fmt.Println(marks[1])
    fmt.Println(marks[5])
    fmt.Println(marks[3])

    name := [5]byte {'f','u','z','z','y'}
    fmt.Printf("%s\n",name)
    name[0] = 'b'
    name[4] = 'z'
    fmt.Printf("%s\n",name)

    cart := [...]string {"Bag", "Shirt", "Watch", "Book"}
    fmt.Printf("There are %d items in your cart\n", len(cart))

    code := [7]rune {'#', '5', 'g', 't', 'm', 'y', '6'}
    fmt.Println("The length of the array is :", len(code))

    for i := 0; i<len(code); i++{
        fmt.Printf("%c\n",code[i])
    }

    for _, s := range cart{
        fmt.Println(s)
    }
}
