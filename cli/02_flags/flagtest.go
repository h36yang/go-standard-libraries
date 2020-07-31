package main

import (
	"flag"
	"fmt"
)

func main() {
	// Create a string flag with default value
	archPtr := flag.String("arch", "x86", "CPU Type")
	// Create a boolean flag with default value
	debugPtr := flag.Bool("debug", false, "Debug Mode")
	// Parse the flags from command
	flag.Parse()

	if *debugPtr {
		fmt.Println("Running in Debug Mode")
	} else {
		fmt.Println("Running in Production")
	}

	switch *archPtr {
	case "x86":
		fmt.Println("Running in 32 bit mode")
	case "AMD64":
		fmt.Println("Running in 64 bit mode")
	case "IA64":
		fmt.Println("Remember IA64?")
	}
	fmt.Println("Process Complete")
}
