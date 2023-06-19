package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("%T\n", now)
	fmt.Println(now)

	fmt.Println(now.UTC())

	somedate := time.Date(2020, 03, 25, 8, 5, 0, 0, time.FixedZone("UTC", 5))
	fmt.Println(somedate)

	today := somedate.Day()
	fmt.Println(today)

	year := somedate.Year()
	fmt.Println(year)

	month := somedate.Month()
	fmt.Println(month)

	date := somedate.Day()
	fmt.Println(date)

	weekday := somedate.Weekday()
	fmt.Println(weekday)

	yearDay := somedate.YearDay()
	fmt.Println(yearDay)


	year, month, date = now.Date()

	fmt.Println(year, month, date)

	fmt.Println(now.Format("Monday 01 January 2006 15:04:05"))

	// day month date hour:minutes:second timezone year
	fmt.Println(now.Format(time.UnixDate))

	// day month date hour:minutes:second year
	fmt.Println(now.Format(time.ANSIC))

	// day month date hour:minutes:second
	fmt.Println(now.Format(time.Stamp))

	// day month date hour:minutes:second.milisecond
	fmt.Println(now.Format(time.StampMilli))

	// day month date hour:minutes:second.microsecond
	fmt.Println(now.Format(time.StampMicro))

	// day month date hour:minutes:second.nanosecond
	fmt.Println(now.Format(time.StampNano))

	// day, date month year hour:minutes:second timezone
	fmt.Println(now.Format(time.RFC1123))

	// day, date month year hour:minutes:second offset
	fmt.Println(now.Format(time.RFC1123Z))

	// year-month-dayThour:minutes:second+-offset
	fmt.Println(now.Format(time.RFC3339))

	// year-month-dayThour:minutes.nanosecond:second
	fmt.Println(now.Format(time.RFC3339Nano))

	// date month year hour:minutes timezone
	fmt.Println(now.Format(time.RFC822))

	// hour:minuteAM/PM
	fmt.Println(now.Format(time.Kitchen))
	customDate := "2023-0426"
	t, err := time.Parse("2006-01-02", customDate)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t)

}
