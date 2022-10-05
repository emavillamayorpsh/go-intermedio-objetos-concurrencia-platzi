package main

import (
	"github.com/donvito/hellomod"
	hellomod2 "github.com/donvito/hellomod/v2" // since both libraries have the same name ("hellomod") , we add the alias "hellomod2" to distinguish them
)

func main() {
	hellomod.SayHello()
	hellomod2.SayHello("course")
}
