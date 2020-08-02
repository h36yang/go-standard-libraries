package main

import (
	"fmt"
	"time"
)

func main() {
	/*
	 * Reference Time:
	 * Mon Jan 2 15:04:05 -0700 MST 2006
	 */
	// Time increments with Sleep()
	first := time.Now()
	fmt.Printf("It is currently %s\n", first.Format("15:04:05"))
	time.Sleep(2 * time.Second)
	second := time.Now()
	fmt.Printf("It is now %s\n", second.Format("15:04:05"))

	// Calculate time span SINCE a date
	today := time.Now()
	fmt.Printf("It is currently %s\n", today.Format("Monday, Jan 2, 2006"))
	startDate := time.Date(2018, 7, 4, 9, 0, 0, 0, time.UTC)
	elapsed := time.Since(startDate)
	fmt.Printf("%v has passed since %s\n", elapsed, startDate.Format("Monday, Jan 2, 2006"))
	fmt.Printf("Hours: %.0f Minutes: %.0f Seconds: %.0f\n", elapsed.Hours(), elapsed.Minutes(), elapsed.Seconds())

	// Calculate past/future dates by adding/subtracting time span
	future := today.AddDate(0, 6, 0)  // increments of year/month/date
	past := today.Add(-6 * time.Hour) // increments of hour/min/sec/...
	fmt.Printf("In 6 months, it will be %s\n", future.Format("Monday, Jan 2, 2006"))
	fmt.Printf("6 hours ago, it was %s\n", past.Format("15:04:05"))

	// Calculate time span UNTIL a date
	deadline := time.Date(2020, 9, 10, 11, 12, 13, 14, time.Local)
	timeLeft := time.Until(deadline)
	fmt.Printf("There are %.0f hours left until deadline\n", timeLeft.Hours())
}
