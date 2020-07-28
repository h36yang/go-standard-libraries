package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Using bufio scanner to scan keyboard inputs
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "/quit" {
			fmt.Println("Quitting")
			os.Exit(3)
		} else {
			fmt.Println("You typed", scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
