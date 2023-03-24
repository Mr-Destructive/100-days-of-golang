package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	// set the seed as the current tim in nano seconds
	// this will set the initial value of the RNG so that
	// each time a new sequence is generated based on the seed value
	rand.Seed(time.Now().UnixNano())

	num := rand.Int()
	log.Println(num)

	for i := 0; i < 30; i++ {
		// generate a integer between 0 and 5
		dice_throw := rand.Intn(6)
		// Move the Offset of 0
		log.Println(dice_throw + 1)

	}

	log.Println(rand.Perm(10))
	arr := []int{1, 2, 3, 4, 5, 6}
	log.Println("Before shuffle:", arr)
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	log.Println("After shuffle:", arr)
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	log.Println("After shuffle:", arr)

	//random characters
	for i := 0; i < 30; i++ {
		// Uppercase character
		random_byte := rand.Intn(26) + 65
		log.Printf("character = %d : %c", random_byte, random_byte)
		// Lowercase character
		random_byte = rand.Intn(26) + 97
		log.Printf("character = %d : %c", random_byte, random_byte)
	}

	letters := "abcdefghijklmnopqrstuvwxyz"
	shuffled := rand.Perm(len(letters))

	result := make([]byte, len(letters))
	for i, randIndex := range shuffled {
		result[i] = letters[randIndex]
	}

	log.Println(string(result))
}
