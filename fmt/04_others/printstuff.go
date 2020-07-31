package main

import "fmt"

type point struct {
	x, y int
}

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	// Print Value
	p := point{x: 2, y: 3}
	fmt.Printf("%v\n", p)

	// Print Type
	newPerson := person{firstName: "Bruce", lastName: "Wayne", age: 28}
	fmt.Printf("%T\n", newPerson)

	// Print Boolean
	var isCool = true
	fmt.Printf("Value is %t\n", isCool)

	// Print Decimals
	fmt.Printf("%d\n", 42)
	fmt.Printf("%b\n", 42) // binary format
	fmt.Printf("%x\n", 42) // hex format
	fmt.Printf("%c\n", 42) // ascii symbol
}
