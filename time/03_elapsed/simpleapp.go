package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	args := os.Args

	// open the customers list
	custlist, err := os.Open(string(args[1]))
	check(err)
	defer custlist.Close()
	writeTime(start, "Opened Customers List")

	// create an output file
	outfile, err := os.Create("outfile.out")
	check(err)
	defer outfile.Close()
	writeTime(start, "Created Output File")

	// scan the customers list
	scanner := bufio.NewScanner(custlist)
	for scanner.Scan() {
		names := strings.Split(scanner.Text(), ",")
		outfile.WriteString(names[1] + "," + names[2] + "\n")
	}
	check(scanner.Err())
	writeTime(start, "Wrote Date to Output File")

	defer writeTime(start, "Application Finished")
}

func writeTime(start time.Time, name string) {
	// Calculate elapsed time with monotonic clock
	elasped := time.Since(start)
	log.Printf("%s took %s\n", name, elasped)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
