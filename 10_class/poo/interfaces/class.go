package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
	endDate string
}

// in order to apply an interface we need to do it in a "implicit" way by creating a function
// with the same name of the interface function/field
func (ftEmployee FullTimeEmployee) getMessage () string {
	return "Full time Employee"
}


type TemporaryEmployee struct {
	Person
	Employee
	taxRate int
}

// apply interface function to TemporaryEmployee class
func (tEmployee TemporaryEmployee) getMessage () string {
	return "Temporary time Employee"
}

// interface creation
type PrintInfo interface {
	getMessage() string
}

// interface definition
func getMessage (p PrintInfo) {
	fmt.Println(p.getMessage())
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Name"
	ftEmployee.age = 2
	ftEmployee.id = 5
	//GetMessage(ftEmployee)

	tEmployee := TemporaryEmployee {}
	getMessage(tEmployee)
	getMessage(ftEmployee)
}
