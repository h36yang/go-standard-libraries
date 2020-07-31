package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 || args[0] == "/help" {
		fmt.Println("Usage: filemaker <Input File>")
	} else {
		fmt.Println("How would you like to see the text?")
		fmt.Println("1: UPPER CASE")
		fmt.Println("2: Title Case")
		fmt.Println("3: lower case")

		// Take user input
		var option int
		_, err := fmt.Scanf("%d", &option)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Open up the file
		file, err := os.Open(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		// Scan file and print each line
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			switch option {
			case 1:
				fmt.Println(strings.ToUpper(scanner.Text()))
			case 2:
				fmt.Println(strings.Title(scanner.Text()))
			case 3:
				fmt.Println(strings.ToLower(scanner.Text()))
			default:
				fmt.Println(scanner.Text())
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}
	}
}
