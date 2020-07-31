package main

import (
	"fmt"
	"os"
)

type messageType int

// Message Type constants
const (
	INFO messageType = iota
	WARNING
	ERROR
)

// Message Color Format constants
const (
	InfoColor    = "\033[1;34m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
)

func main() {
	fileName := "test.txt"
	showMessage(INFO, fmt.Sprintf("About to open %s", fileName))
	showMessage(WARNING, "If the file is not present, the application will fail.")

	file, err := os.Open(fileName)
	if err != nil {
		showMessage(ERROR, err.Error())
	}
	defer file.Close()
}

// Print message in different format based on type
func showMessage(msgType messageType, message string) {
	switch msgType {
	case INFO:
		printMessage := fmt.Sprintf("\nInformation: \n%s\n", message)
		fmt.Printf(InfoColor, printMessage)
	case WARNING:
		printMessage := fmt.Sprintf("\nWarning: \n%s\n", message)
		fmt.Printf(WarningColor, printMessage)
	case ERROR:
		printMessage := fmt.Sprintf("\nError: \n%s\n", message)
		fmt.Printf(ErrorColor, printMessage)
	}
}
