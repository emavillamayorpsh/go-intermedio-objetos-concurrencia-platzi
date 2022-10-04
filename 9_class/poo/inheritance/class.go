package main

import "fmt"

type Person struct {
	name string
	age int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	// anonymous field
	Person
	Employee
}

func GetMessage (p FullTimeEmployee) {
	fmt.Printf("%s with age %d\n" , p.name , p.age)
}

func main () {
	// this instance has properties of "Person" and "Employee" classes
	ftEmployee := FullTimeEmployee{}

	ftEmployee.name = "Name" // field from Person class
	ftEmployee.age = 2 // field from Person class
	ftEmployee.id = 5 // field from Employee class

	fmt.Println("ftEmployee: ", ftEmployee)
	// the printed object looks like this:
	// {{Name 2} {5}}
	// in the first object there are the fields from "Person" class (name and age) , and
	// in the second object there is the field from Employee class (id)


	GetMessage(ftEmployee)
}