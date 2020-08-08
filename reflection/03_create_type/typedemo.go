package main

import (
	"fmt"
	"reflect"
)

type employee struct {
	id        int
	firstName string
	lastName  string
}

func main() {
	employees := make([]employee, 0)
	employees = append(employees, employee{0, "Bruce", "Wayne"})
	employees = append(employees, employee{1, "Clark", "Kent"})
	employees = append(employees, employee{2, "Diana", "Prince"})

	// Create new slice of employees using reflection
	eType := reflect.TypeOf(employees)
	newEmployeeList := reflect.MakeSlice(eType, 0, 0)
	newEmployeeList = reflect.Append(newEmployeeList, reflect.ValueOf(employees[0]))
	newEmployeeList = reflect.Append(newEmployeeList, reflect.ValueOf(employees[1]))
	newEmployeeList = reflect.Append(newEmployeeList, reflect.ValueOf(employees[2]))

	fmt.Printf("First list of employees: %v\n", employees)
	fmt.Printf("List created by reflection: %v\n", newEmployeeList)
}
