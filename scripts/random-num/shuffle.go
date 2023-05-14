package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	fmt.Println(r.Perm(10))
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("Before shuffle:", arr)
	r.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	fmt.Println("After shuffle:", arr)
	r.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	fmt.Println("After shuffle:", arr)

	//random characters
	for i := 0; i < 30; i++ {
		// Uppercase character
		random_byte := r.Intn(26) + 65
		fmt.Printf("character = %d : %c\n", random_byte, random_byte)
		// Lowercase character
		random_byte = r.Intn(26) + 97
		fmt.Printf("character = %d : %c\n", random_byte, random_byte)
	}

	//Random string
	randomLowerCase := make([]rune, 6)
	randomUpperCase := make([]rune, 6)
	fmt.Println(randomLowerCase)
	fmt.Println(randomUpperCase)
	for i := range randomLowerCase {
		randomLowerCase[i] = rune(r.Intn(26) + 97)
		randomUpperCase[i] = rune(r.Intn(26) + 65)
	}
	randomLowerCaseStr := string(randomLowerCase)
	randomUpperCaseStr := string(randomUpperCase)
	fmt.Println(randomLowerCase)
	fmt.Println(randomLowerCaseStr)
	fmt.Println(randomUpperCase)
	fmt.Println(randomUpperCaseStr)

	//Randomly shuffled alphabets
	letters := "abcdefghijklmnopqrstuvwxyz"
	shuffled := r.Perm(len(letters))

	result := make([]byte, len(letters))
	for i, randIndex := range shuffled {
		result[i] = letters[randIndex]
	}

	rand_str := string(result)
	fmt.Println(rand_str)
	fmt.Println(rand_str[:10])
}
