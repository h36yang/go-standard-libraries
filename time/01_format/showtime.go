package main

import (
	"fmt"
	"time"
)

// Defining custom formats as constants
const (
	layoutISO = "2006-01-02"
	layoutUS  = "January 2, 2006"
	layoutEU  = "2 January, 2006"
)

func main() {
	// Get current date/time
	t := time.Now()
	fmt.Println(t)

	// Extracting Year/Month/Day separately
	fmt.Println(t.Year())
	fmt.Println(t.Month())
	fmt.Println(t.Day())

	// Built-in formats
	fmt.Println(t.Format(time.ANSIC))
	fmt.Println(t.Format(time.RFC3339))
	fmt.Println(t.Format(time.UnixDate))
	fmt.Println(t.Format(time.Kitchen))

	// Injecting own format:
	// Format value must be 'Mon Jan 2 15:04:05 -0700 MST 2006' to work
	// Refer to https://golang.org/pkg/time/#Time.Format
	fmt.Println(t.Format("Monday, January 2 in the year of 2006"))

	// Create and format a custom date
	startDate := time.Date(2018, 7, 4, 9, 0, 0, 0, time.UTC)
	fmt.Println(startDate)
	fmt.Println(startDate.Format(layoutISO))
	fmt.Println(startDate.Format(layoutUS))
	fmt.Println(startDate.Format(layoutEU))
}
