package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int // explicit declaration
	x = 8
	y := 7 // implicit declaration

	fmt.Println(x)
	fmt.Println(y)

	// returns int value and error
	myValue, err := strconv.ParseInt("7s", 0 , 64)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("My value: ", myValue)
	}

	// create a map with string KEYS and int VALUES
	m := make(map[string]int)
	m["key"] = 6
	fmt.Println("map key: ", m["key"])

	// create an slice of int and set
	s := [] int{1,2,3}
	for index, value := range s {
		fmt.Println("index: ",index)
		fmt.Println("value: ", value)
	}

	// add new value at end of slice
	s = append(s, 16)

	// post add new value
	for index, value := range s {
		fmt.Println("index: ",index)
		fmt.Println("value: ", value)
	}

}