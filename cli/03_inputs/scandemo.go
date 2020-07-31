package main

import "fmt"

func main() {
	var name string
	fmt.Println("What is your name?")
	// Use fmt package to scan input into &name pointer
	inputs, _ := fmt.Scanf("%q", &name)
	switch inputs {
	case 0:
		fmt.Println("You must enter a name!")
	case 1:
		fmt.Printf("Hello %s! Nice to meeting you\n", name)
	}
}
