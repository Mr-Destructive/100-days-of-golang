package main

import (
	"fmt"
	"time"
)

func main() {
	screentime, err := time.ParseDuration("6h30m")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T\n", screentime)
	fmt.Println(screentime.Hours())
	fmt.Println(screentime.Minutes())

	newYear := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println(time.Now().UTC())
	fmt.Println(time.Since(newYear).Hours())

	nextNewYear := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println(nextNewYear.Sub(newYear).Hours())
	fmt.Println(nextNewYear.Sub(newYear).String())

	day := time.Hour * 24
	fmt.Println(day)
	week := time.Hour * 24 * 7
	fmt.Println(week)
	month := time.Hour * 30 * 24 * 7
	fmt.Println(month)

	fifteenDays := day * 15
	fmt.Println(fifteenDays)
}
