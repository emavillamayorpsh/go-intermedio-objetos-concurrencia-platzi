package main

import (
	"fmt"
	"time"
)

func main() {
	// we created a channel that will monitor the "doSomething" function
	c :=  make(chan int)

	// The reserved word "go" allows us to create a "go routine" of "doSomething" function
	// this means that the function will be executed in a routine different of "main"
	go doSomething(c)

	// with the following code main will stop it's execution until
	// the channel return the value "1" (that happens after three seconds  *code in function doSomething*)
	<-c


	// POINTERS
	g:= 25
	fmt.Println(g)
	h := &g // memory address / reference where "g" is located
	fmt.Println(h)
	fmt.Println(*h) // value of the memory address
}

// receives a channel as param
func doSomething(c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Done")

	// it will return the value of 1
	c <- 1
}