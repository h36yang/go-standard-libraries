package main

import "fmt"

func main() {
	var name string = "Jack"
	var age int = 18

	// Use fmt.Print() to print tokenized strings
	fmt.Print("My name is ", name, " and I'm ", age, " years old.\n")

	// Use fmt.Printf() to print formatted strings with %v
	fmt.Printf("My name is %v and I'm %v years old.\n", name, age)

	// Use fmt.Printf() with implicit data type specifiers
	fmt.Printf("My name is %s and I'm %d years old.\n", name, age)

	/*
	 * Common Format Specifiers:
	 * 	%v - formats the value in a default format
	 * 	%s - formats string values
	 *	%q - formats string values quoted
	 * 	%d - formats decimal integers
	 * 	%g - formats floating-point numbers. %e for large exponents, %f otherwise.
	 *	%e - formats floating-point numbers in scientific notation, e.g. -1.234456e+78
	 *	%f - formats floating-point numbers but no exponent, e.g. 123.456
	 * 	%b - formats base-2 numbers
	 *	%o - formats base-8 numbers
	 *	%t - formats true or false values
	 */

	// Use fmt.Printf() to format floats
	var pi float32 = 3.141592
	fmt.Printf("Pi is %f\n", pi)
	fmt.Printf("Pi is %.2f\n", pi)

	/*
	 * Use fmt.Printf() to print nicely formatted tables
	 */

	// Floats
	fmt.Printf("|%7.2f|%7.2f|%7.2f|\n", 23.3774, 577.45, 1234.56)
	fmt.Printf("|%7.2f|%7.2f|%7.2f|\n", 3287.921737, 23.324, 98.99)

	// Left Aligned
	fmt.Printf("|%-7.2f|%-7.2f|%-7.2f|\n", 23.3774, 577.45, 1234.56)
	fmt.Printf("|%-7.2f|%-7.2f|%-7.2f|\n", 3287.921737, 23.324, 98.99)

	// Strings
	fmt.Printf("|%-7s|%-7s|%-7s|\n", "foo", "bar", "go")
	fmt.Printf("|%-7s|%-7s|%-7s|\n", "a", "ab", "abc")

	// Quoted Strings
	fmt.Printf("|%-7q|%-7q|%-7q|\n", "foo", "bar", "go")
	fmt.Printf("|%-7q|%-7q|%-7q|\n", "a", "ab", "abc")

	// Use fmt.Sprintf() to save formatted output and print later
	output := fmt.Sprintf("|%-7q|%-7q|%-7q|\n", "a", "ab", "abc")
	print(output)
}
