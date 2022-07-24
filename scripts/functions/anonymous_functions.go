package main

import "fmt"
import "strings"

func get_csv(index int, str string, t func(csv string)[]string) [][]string{
  s := t(str)
  var res [][]string
  for i := 1; i<len(s); i+=2 {
    var data []string
    data = append(data, s[i-1], s[i])
    res = append(res, data)
  }
  return res
}

func main() {

  // Simple Anonymous function
  func() {
    fmt.Println("Hello, Anonymous function")
  }()

  // Assinging the Anonymous function to a variable
  draw := func() {
    fmt.Println("Drawing")
  }
  draw()

  // passing parmaters to the Anonymous function
  call := func(name string) {
    fmt.Println("Calling,", name)
  }
  call("Meet")
  person := "Chris"
  call(person)

  // Returning values from Anonymous function
  is_divisible := func(x int, y int) bool{
    var res bool
    if x % y == 0 {
      res = true
    } else{
      res = false
    }
    fmt.Println(res)
    return res
  }

  fmt.Printf("%T\n",is_divisible)
  is_divisible(10, 5)
  is_divisible(33, 5)

  divisblity_check := is_divisible(45, 5)
  fmt.Printf("%T : %t\n", divisblity_check, divisblity_check) 

  format_email := func(name string, age int, company string) string{
    email := fmt.Sprintf("%s.%d@%s.com", name, age, company)
    return email
  }("john", 25, "gophersoft")

  fmt.Println(format_email)
  fmt.Printf("%T\n",format_email)

  // Passing function literal to a function parmater
  csv_slice := func (csv string) []string{
    return strings.Split(csv, ",")
  }
  csv_data := get_csv(2,"kevin,21,john,33,george,24", csv_slice)
  fmt.Println(csv_data)
  for _, s := range csv_data{
    fmt.Printf("%s - %s\n",s[0],s[1])
  }
}
