package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Read arguments from os package
	args := os.Args[1:]
	if len(args) == 1 && args[0] == "/help" {
		fmt.Println("Usage: dinnertotal <Total Mean Amount> <Tip Percentage>")
		fmt.Println("Example: dinnertotal 20 10")
	} else {
		if len(args) != 2 {
			fmt.Println("You must enter 2 arguments! Type /help for help")
		} else {
			mealTotal, _ := strconv.ParseFloat(args[0], 32)
			tipAmount, _ := strconv.ParseFloat(args[1], 32)
			fmt.Printf("Your meal total will be $%.2f\n", calculateTotal(float32(mealTotal), float32(tipAmount)))
		}
	}
}

func calculateTotal(mealTotal, tipAmount float32) float32 {
	totalPrice := mealTotal + mealTotal*tipAmount/100
	return totalPrice
}
