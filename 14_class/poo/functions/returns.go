package main

import "fmt"

// "..." treats the "values" prop as an slice ([])
// the difference of doing (values [] int) is that when implementing the function we pass different params (1,2,3,4)
// instead of an "slice" ([1,2,3,4])
func sum(values ...int) int {
	total := 0
	for _, num := range values {
		total += num
	}

	return total
}

func printNames (names ...string) {
	for _, name := range names {
		fmt.Printf("%s ", name)
	}
}

// return values with names
// if we define multiple values to return we can name them and use them inside the function by passing it's values
func getValues (x int ) (double int, triple int, quad int) {
	double = 2 * x
	triple = 3 * x
	quad = 4 * x

	// we don't need to return every single return param because those are already populated above
	return
}

func main() {
	fmt.Println(sum(1))
	fmt.Println(sum(1,2))
	fmt.Println(sum(1,2,3,4))

	printNames("Manu", "Ale")

	fmt.Println(getValues(1))
}