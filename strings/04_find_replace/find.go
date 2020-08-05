package main

import (
	"fmt"
	"strings"
)

func main() {
	sampleString := "I really enjoy the Go language. It's one of my favourites"

	// Find String Contains
	searchTerm := "Go"
	result := strings.Contains(sampleString, searchTerm)
	fmt.Printf("The sample text includes %q: %t\n", searchTerm, result)

	// Find String Starts With
	searchTerm2 := "I"
	result2 := strings.HasPrefix(sampleString, searchTerm2)
	fmt.Printf("The sample text starts with %q: %t\n", searchTerm2, result2)

	// Find String Ends With
	searchTerm3 := "favourites"
	result3 := strings.HasSuffix(sampleString, searchTerm3)
	fmt.Printf("The sample text ends with %q: %t\n", searchTerm3, result3)

	// Replace String
	sampleString = "This is my string. There are many strings like it, but this one is mine."
	sampleString = strings.Replace(sampleString, "string", "compiler", -1)
	fmt.Println(sampleString)

	// Replace String (up to N times)
	sampleString = "This is my string. There are many strings like it, but this one is mine."
	sampleString = strings.Replace(sampleString, "string", "compiler", 1) // Only replace the first occurrence
	fmt.Println(sampleString)
}
