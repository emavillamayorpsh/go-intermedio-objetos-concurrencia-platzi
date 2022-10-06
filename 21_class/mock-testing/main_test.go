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
			expectedEmployee: FullTimeEmployee{
				Person: Person {
					Name: "John Doe",
					Age: 35,
					DNI: "1",
				},
				Employee: Employee{
					Id: 1,
					Position: "CEO",
				},
			},
		},
	}

	// save the functions behavior without the mock applied
	originalGetEmployeeById := GetEmployeeById
	originalGetPersonByDNI := GetPersonByDNI

	for _, item := range table {
		// the first thing that i do is call the mock in order to override functions:
		// GetEmployeeById and GetPersonByDNI
		item.mockFunc()
		ft, err := GetFullTimeEmployeeById(item.id, item.dni)

		if err != nil {
			t.Errorf("Error when getting Employee")
		}

		if ft.Age != item.expectedEmployee.Age {
			t.Errorf("Error, got %d expected %d", ft.Age, item.expectedEmployee.Age)
		}
	}

	// set back to default the functions (now they don't have the mock applied)
	GetEmployeeById = originalGetEmployeeById
	GetPersonByDNI = originalGetPersonByDNI
}