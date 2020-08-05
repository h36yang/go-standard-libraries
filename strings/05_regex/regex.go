package main

import (
	"fmt"
	"regexp"
)

func main() {
	sampleString := "This is my song. There are many strings like it, but this one is mine."

	r, _ := regexp.Compile(`s([a-z]+)g`) // find all char sequences that start with 's' and end with 'g'
	fmt.Println(r.MatchString(sampleString))
	fmt.Println(r.FindAllString(sampleString, -1))
	fmt.Println(r.FindAllStringIndex(sampleString, -1))
	newString := r.ReplaceAllString(sampleString, "apple")
	fmt.Println(newString)

	r2, _ := regexp.Compile(`s(\w[a-z]+)g\b`) // find all words that start with 's' and end with 'g'
	fmt.Println(r2.MatchString(sampleString))
	fmt.Println(r2.FindAllString(sampleString, -1))
	fmt.Println(r2.FindAllStringIndex(sampleString, -1))
	newString2 := r2.ReplaceAllString(sampleString, "apple")
	fmt.Println(newString2)
}
