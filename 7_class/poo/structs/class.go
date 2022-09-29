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

// Alternative to add functions to structs (similar to class methods)

// We specify that "Employee" struct will have a method called SetId
// that means if we have an instance of "Employee" we will be able to call this function
func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) GetId() int {
	return e.id
}

func (e *Employee) GetName() string {
	return e.name
}

func main() {
	// create instance of Employee
	e := Employee{}
	fmt.Println("Employee: ", e)
	e.id = 1
	e.name = "Name"
	fmt.Println("Employee post assign values: ", e)

	// using struct methods
	e.SetId(2)
	e.SetName("Name 2")
	fmt.Println("Employee post assign with method value: ", e)

	fmt.Println("GetId: ", e.GetId())
	fmt.Println("GetName: ", e.GetName())

}