package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

func NewEmployee (id int , name string,  vacation bool) *Employee {
	// "&" will allows us to modify the values and we need to return the reference
	return &Employee{
		id: id,
		name: name,
		vacation: vacation,
	}
}

func main() {
	// 1
	e := Employee{}
	fmt.Printf("%v\n", e)

	// 2
	e2 := Employee {
		id: 1,
		name: "Name",
		vacation: true,
	}
	fmt.Println("e2: ", e2)

	// 3
	e3 := new(Employee) // this returns a "reference" (&)
	fmt.Println("e3: ", *e3) // add "*" symbol because e3 is a "reference" and we want the "value"
	e3.id = 2
	e3.name = "Name 2"
	fmt.Println("e3 assigned values: ", *e3)

	// 4
	e4 := NewEmployee(3, "Name 3", true)
	fmt.Println("e4: ", e4)

}
