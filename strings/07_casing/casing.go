package main

import (
	"fmt"
	"strings"
)

func main() {
	sampleString := "Never trust a programmer who carries a screwdriver"
	fmt.Println("Before:", sampleString)
	fmt.Println("Lower Case:", strings.ToLower(sampleString))
	fmt.Println("Upper Case:", strings.ToUpper(sampleString))
	fmt.Println("Title Case:", strings.Title(sampleString))
	fmt.Println("Proper Title Case:", properTitle(sampleString))
}

func properTitle(input string) string {
	words := strings.Fields(input)
	smallwords := " a an on the to "

	for index, word := range words {
		if !strings.Contains(smallwords, " "+word+" ") {
			words[index] = strings.Title(word)
		}
	}
	return strings.Join(words, " ")
}
