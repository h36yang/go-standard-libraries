package main

import (
	"fmt"
	"strings"
)

func main() {
	// Trim spaces
	sampleString := "          This is our text           "
	fmt.Printf("%q\n", sampleString)
	newString := strings.TrimSpace(sampleString)
	fmt.Printf("%q\n", newString)

	// Trim prefix
	site := "https://www.google.com"
	domainName := strings.TrimPrefix(site, "https://")
	fmt.Println(domainName)

	// Trim suffix
	fileName := "test.doc"
	fileNameWithoutExt := strings.TrimSuffix(fileName, ".doc")
	fmt.Println(fileNameWithoutExt)
}
