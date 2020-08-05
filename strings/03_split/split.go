package main

import (
	"fmt"
	"strings"
)

func main() {
	// Split the string by separator
	outString := "This is a string!|This is another one|I like turtles"
	stringCollection := strings.Split(outString, "|")
	for i := range stringCollection {
		fmt.Println(stringCollection[i])
	}

	// Split the string by separator (including the separator)
	stringCollection = strings.SplitAfter(outString, "|") // This will include the "|"
	for i := range stringCollection {
		fmt.Println(stringCollection[i])
	}

	// Split the string by separator (up to N splits)
	stringCollection = strings.SplitN(outString, "|", 2) // This will return up to 2 substrings only
	for i := range stringCollection {
		fmt.Println(stringCollection[i])
	}
}
