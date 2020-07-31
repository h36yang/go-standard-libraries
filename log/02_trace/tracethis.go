package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime/trace"
	"time"
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("We did not create a trace file! %v\n", err)
	}
	// defer statements are LIFO
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("Failed to close trace file! %v\n", err)
		}
	}() // this defer will be executed last in the end

	if err := trace.Start(f); err != nil {
		log.Fatalf("Failed to start a trace! %v\n", err)
	}
	defer trace.Stop() // this defer will be executed first in the end

	addRandomNumbers()
}

func addRandomNumbers() {
	firstNumber := rand.Intn(100)
	secondNumber := rand.Intn(100)
	time.Sleep(2 * time.Second)
	result := firstNumber * secondNumber
	fmt.Printf("Result of 2 random numbers is %d\n", result)
}
