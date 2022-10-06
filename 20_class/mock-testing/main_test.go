package main

import "testing"

func TestGetFullTimeEmployeeById(t *testing.T) {
	table := []struct {
		id int
		dni string
		mockFunc func()
		expectedEmployee FullTimeEmployee
	}{
		{
			id: 1,
			dni: "1",
			mockFunc: func() {
				// Override the function "GetEmployeeById" that it is call inside of "GetFullTimeEmployeeById"
				// in order to avoid a call to the backend and return the value that we want
				GetEmployeeById = func (id int)(Employee, error)  {
					return Employee{
						Id: 1,
						Position: "CEO",
					}, nil
				}
				GetPersonByDNI = func(id string) (Person, error) {
					return Person{
						Name: "John Doe",
						Age: 35,
						DNI: "1",
					}, nil
				}
			},
		},
	}
}