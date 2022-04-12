package main

import "fmt"

func main() {

	// simple map litetal
	char_freq := map[string]int{
		"M": 1,
		"e": 2,
		"t": 1,
	}
	fmt.Println(char_freq)

	// declare map with make function
	marks := make(map[int]int)
	fmt.Println(len(marks))
	marks[65] = 8
	marks[95] = 3
	marks[80] = 5
	fmt.Println(marks)

	// declare map with new function
	name := new(map[byte]int)
	*name = map[byte]int{}
	name_map := *name
	name_map['m'] = 1
	name_map['e'] = 2
	name_map['t'] = 1
	fmt.Println(len(name_map))
	fmt.Println(name_map)

	// check if a key exist in map
	var key byte = 't'
	value, exist := name_map[key]
	if exist == true {
		fmt.Printf("The key %c exist and has value %d\n", key, value)
	} else {
		fmt.Printf("The key %c does not exist.\n", key)
	}

	for k, _ := range name_map {
		fmt.Printf("%c\n", k)
	}

	cart_list := map[string]int{
		"shirt": 2,
		"mug":   4,
		"shoes": 3,
	}
	fmt.Println(cart_list)

	// add new key in map
	cart_list["jeans"] = 1

	// update existing keys in map
	cart_list["mug"] = 3

	// delete keys from map
	delete(cart_list, "shoes")
	fmt.Println(cart_list)

	is_prime := map[int]bool{
		7:  true,
		9:  false,
		13: true,
		15: false,
		16: false,
	}

	// iterate over map
	for key, value := range is_prime {
		fmt.printf("%d -> %t\n", key, value)
	}

	// iterate only with key or value
	for key, _ := range is_prime {
		fmt.Printf("Key : %t\n", _)
	}

	for _, value := range is_prime {
		fmt.Printf("Value: %t\n", value)
	}

}
