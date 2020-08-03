package main

import (
	"fmt"
	"strings"
)

func main() {
	// strings.Compare()
	string1 := "This is a string!"
	string2 := "this is a string!"
	if strings.Compare(string1, string2) == 0 {
		fmt.Println("The strings are identical!")
	} else {
		fmt.Println("The strings do not match!")
	}

	stooges := []string{"Larry", "Curly", "Moe"}
	for _, stooge := range stooges {
		fmt.Println(strings.Compare("Larry", stooge))
	}

	// Compare case insensitive
	fmt.Println(compareCaseIns(string1, string2))
}

func compareCaseIns(a, b string) bool {
	if len(a) == len(b) {
		if strings.ToLower(a) == strings.ToLower(b) {
			return true
		}
	}
	return false
}
