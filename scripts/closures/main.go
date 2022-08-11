package main

import "fmt"

func gophy() func() string{
  return func() string{
    return "Hello, Gophers!"
  }
}

func inrementer() func() int{
  counter := 0 
  return func() int{
    counter += 1
    return counter
  }
}

func factorial() func() int{
	fact, n := 1, 1
	return func() int{
    fact = fact * n
    n += 1
		return fact
	}
}

func main() {
  // s := "Hello, Gophers!"
  // long form of declaring s
  // using clousure/anonymous function to return a value
  // that value can be assigned to the variable
  s := func() string {
    return "Hello, World!" 
  }()
  fmt.Println(s)
  g := gophy()
  fmt.Println(g())

  c := inrementer()
  fmt.Println(c())
  fmt.Println(c())
  fmt.Println(c())

  c1 := inrementer()
  fmt.Println(c1())
  fmt.Println(c1())
  fmt.Println(c1())


  f1 := factorial()
  fmt.Println(f1())
  fmt.Println(f1())
  fmt.Println(f1())
}
