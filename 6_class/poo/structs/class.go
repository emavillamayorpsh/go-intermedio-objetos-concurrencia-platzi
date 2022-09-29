package main

import "fmt"

// structs are equivalent to classes
// they will allows us to define some "properties"
// meaning only the "attributes" or "properties" of a "class"

type Employee struct {
	// Each property have default values in case they are not define whe instancing an Employee
	id int
	name string
}

func main() {
	// create instance of Employee
	e := Employee{}
	fmt.Println("Employee: ", e)
	e.id = 1
	e.name = "Name"
	fmt.Println("Employee post assign values: ", e)

}