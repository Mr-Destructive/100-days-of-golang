package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Location())

	newYorkTimeZone, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println(err)
	}
	londonTimeZone, err := time.LoadLocation("Europe/London")
	if err != nil {
		fmt.Println(err)
	}
	newYorkTime := t.In(newYorkTimeZone)
	londonTime := t.In(londonTimeZone)

	//local time
	fmt.Println(t)

	// london time
	fmt.Println(londonTimeZone)
	fmt.Println(londonTime)

	// new york time
	fmt.Println(newYorkTimeZone)
	fmt.Println(newYorkTime)

	offset := int((4*time.Hour + 30*time.Minute + 30*time.Second).Seconds())
	// (4*60 + 30*1 + 30*0.166) * 60
	// (270 + 0.5) * 60 = 16230
	fmt.Println(offset)
	fmt.Println(t.UTC())
	customTimeZone := time.FixedZone("SOMETZ", offset)
	fmt.Println(customTimeZone)
	fmt.Println(t.In(customTimeZone))
}
