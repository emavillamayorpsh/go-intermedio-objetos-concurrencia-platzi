package main

import (
	"fmt"
	"time"
)

// anonymous function , the main purpose is when we know that the function is going to be used ONLY ONCE
// when we want to use go routines we can create an anonymous function
func main() {
	x := 5
	y := func() int {
		return x * 2
	}()

	fmt.Println(y)

	// =====================
	// Go routine example
	c := make(chan int)
	go func () {
		fmt.Println("Starting function")
		time.Sleep(2 * time.Second)
		fmt.Println("End")
		c <- 1
	}()
	<-c

}