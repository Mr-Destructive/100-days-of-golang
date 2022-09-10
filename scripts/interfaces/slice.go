package main

import (
	"fmt"
)

type Device interface {
	wifi() string
}

type Laptop struct {
	name string
}

type Phone struct {
	name  string
	price int
	ram   int
}

func (l Laptop) wifi() string {
	fmt.Println("+Connect to wifi with Laptop interface")
	return l.name
}

func (m Phone) wifi() string {
	fmt.Println("+Connect to wifi with Mobile interface")
	return m.name
}

func main() {

	devices := []Device{Laptop{}, Phone{}, Phone{}, Laptop{}}

	for _, obj := range devices {
		fmt.Println(obj.wifi())
	}
}
