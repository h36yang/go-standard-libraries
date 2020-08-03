package main

import "fmt"

func main() {
	// String as slice of bytes
	var outString string = "\x47\x6f\x20\x69\x73\x20\x41\x77\x65\x73\x6f\x6d\x65\x21"
	fmt.Println(outString)
	for i := 0; i < len(outString); i++ {
		// Print the byte hex
		fmt.Printf("%x : ", outString[i])
		// Print the quoted char
		fmt.Printf("%q\n", outString[i])
	}

	// Debug non-printable chars
	outString = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	fmt.Println(outString)
	for i := 0; i < len(outString); i++ {
		// Print the byte hex
		fmt.Printf("%x : ", outString[i])
		// Print the quoted char
		fmt.Printf("%q\n", outString[i])
	}

	// Extract chars from string by index
	newString := "This is a string!"
	fmt.Println(newString)
	fmt.Println(newString[3])   // single index to extract the byte
	fmt.Println(newString[3:4]) // slice index range to extract the char
}
