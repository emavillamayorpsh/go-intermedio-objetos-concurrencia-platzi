package main

import "testing"

// In order to run a test we need to have a module: go mod init github.com/emavillamayorpsh/test
// To run test we ran the following command in the terminal: go test

// testing.T we use this to verify that our test is running correctly
// is native of go
func TestSumBasic(t *testing.T) {
	total := Sum(5, 5)

	if total != 10 {
		// we indicate to go that the test failed , and go will know in which line and why is happening failed
		t.Errorf("Sum was incorrect, got %d expected %d", total , 10)
	}
}

func TestSumComplex(t *testing.T) {
	// create different test scenarios
	tables := []struct {
		a int
		b int
		n int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 26, 51},
	}

	for _, item := range tables {
		total := Sum(item.a , item.b)

		if total != item.n {
			t.Errorf("Sum was incorrect, got %d expected %d", total , item.n)
		}
	}
}

func TestGetMax(t *testing.T) {
	tables := []struct{
		a int
		b int
		n int
	} {
		{ 4, 2, 4 },
		{ 3, 2, 3},
		{2, 5, 5},
	}

	for _, item := range tables {
		max := GetMax(item.a, item.b)

		if max != item.n {
			t.Errorf("GetMax was incorrect, got %d expected %d", max , item.n)
		}
	}
}

// For test that take too much time to render we run the followings commands:
// This command generates a file:  go test -cpuprofile=cpu.out
// Once the file is created run the command: go tool pprof cpu.out
	// * command "top": tells us the behaviour in terms of time , percentage of the test
	// * command "list *nameofthefunction*": tells us how much time it takes between lines of the function
	// * command "web": generates a file svg that it's easier to read
	// * command "pdf:" exports the report
func TestFib(t *testing.T) {
	tables := [] struct {
		a int
		n int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}

	for _, item :=  range tables {
		fib := Fibonacci(item.a)

		if fib != item.n {
			t.Errorf("Fibonacci was incorrect, got %d expected %d", fib , item.n)
		}
	}
}