package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Use fmt.Sscanf() to scan variables from a line of text
	file, err := os.Open("family.csv")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var name string
		var age int
		n, err := fmt.Sscanf(scanner.Text(), "%s is %d years old", &name, &age)
		if err != nil {
			panic(err)
		}

		if n == 2 {
			fmt.Println(name, age)
		}
	}

	// Use fmt.Scanln() to scan the first non-space input
	fmt.Println("What is your name?")
	var nameInput string
	fmt.Scanln(&nameInput)
	fmt.Printf("Hello %s. Nice to meet you!\n", nameInput)
}
