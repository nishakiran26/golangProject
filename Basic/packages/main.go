package main

import (
	"golangProject/packages/helpers"
	"log"
)

func main() {
	log.Println("Hello")

	//var myVar helpers.SomeType
	var myVar helpers.SomeType
	myVar.TypeName = "Some Name"

	// myVar.TypeName = "Some Name"
	log.Println(myVar.TypeName)
}
