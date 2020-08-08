package main

import (
	"fmt"
	"reflect"
)

func main() {
	type employee struct {
		id        int
		firstName string
		lastName  string
	}

	type customer struct {
		id        int
		firstName string
		lastName  string
		company   string
	}

	newEmployee := employee{0, "Clark", "Kent"}
	newCustomer := customer{234, "Bruce", "Wayne", "Wayne Enterprise"}
	fmt.Printf("Our employee is %s %s with an ID of %d\n", newEmployee.firstName, newEmployee.lastName, newEmployee.id)
	fmt.Printf("The type is %v\n", reflect.TypeOf(newEmployee))                     // type
	fmt.Printf("The value is %v\n", reflect.ValueOf(newEmployee))                   // value
	fmt.Printf("The kind of type it is: %v\n", reflect.ValueOf(newEmployee).Kind()) // kind

	addPerson(newEmployee)
	addPerson(newCustomer)
}

// Take a generic interface{} type parameter
func addPerson(p interface{}) bool {
	// Different actions based on kind and type
	v := reflect.ValueOf(p)
	if v.Kind() == reflect.Struct {
		switch reflect.TypeOf(p).Name() {
		case "employee":
			empSQLString := "INSERT INTO employees (id, firstName, lastName) VALUES (?, ?, ?)"
			fmt.Printf("SQL: %s\n", empSQLString)
			fmt.Printf("Added: %v\n", v.Field(1))
		case "customer":
			custSQLString := "INSERT INTO customers (id, firstName, lastName, company) VALUES (?, ?, ?, ?)"
			fmt.Printf("SQL: %s\n", custSQLString)
			fmt.Printf("Added: %v\n", v.Field(1))
		}
		return true
	}
	return false
}
