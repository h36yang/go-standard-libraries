package main

import (
	"log"
	"os"
)

type messageType int

// Message Type constants
const (
	INFO messageType = iota
	WARNING
	ERROR
	FATAL
)

func main() {
	writeLog(INFO, "This is an information message!")
	writeLog(WARNING, "This is a warning message!")
	writeLog(ERROR, "This is an error message!")
	writeLog(FATAL, "This is a fatal message!")
}

func writeLog(msgType messageType, message string) {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	switch msgType {
	case INFO:
		logger := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Println(message)
	case WARNING:
		logger := log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Println(message)
	case ERROR:
		logger := log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Println(message)
	case FATAL:
		logger := log.New(file, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Fatal(message)
	}
}
