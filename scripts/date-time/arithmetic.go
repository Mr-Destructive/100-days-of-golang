package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)

	afterOneHour := t.Add(time.Hour * 1)
	fmt.Println(afterOneHour)

	afterOneDayTwoHours30minutes := t.AddDate(0, 0, 1).Add(time.Hour * 3).Add(time.Minute * 30)
	fmt.Println(afterOneDayTwoHours30minutes)

	afterThreeMonths15days := t.AddDate(0, 3, 15)
	fmt.Println(afterThreeMonths15days)

	beforeOneHour := t.Add(time.Hour * -1)
	fmt.Println(beforeOneHour)

	beforeOneYearTwoMonths := t.AddDate(-1, -2, 0)
	fmt.Println(beforeOneYearTwoMonths)

	isNowAfter := t.After(afterOneHour)
	isOneAfter := afterOneHour.After(t)
	fmt.Println(isNowAfter)
	fmt.Println(isOneAfter)

	isNowBefore := t.Before(afterOneHour)
	isOneBefore := afterOneHour.Before(t)
	fmt.Println(isNowBefore)
	fmt.Println(isOneBefore)

	isNowEqual := t.Equal(afterOneHour)
	isEqual := time.Now().Equal(t)
	fmt.Println(isNowEqual)
	fmt.Println(isEqual)

	londonTimeZone, err := time.LoadLocation("Europe/London")
	if err != nil {
		fmt.Println(err)
	}
	londonTime := t.In(londonTimeZone)

	fmt.Println(t)
	fmt.Println(londonTime)
	fmt.Println(t.Equal(londonTime))
}
