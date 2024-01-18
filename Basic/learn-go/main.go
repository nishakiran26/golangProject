package main

import "fmt"

//var myName string //package level variable
func main() {
	fmt.Println("Hello world")

	var whatToSay string
	var i int

	whatToSay = "Goodbye, cruel world"
	fmt.Println(whatToSay)

	i = 20
	fmt.Println("i is set to : ", i)
	whatWasSaid := saySomething()
	fmt.Println("the function return", whatWasSaid)
	a, b := saySomething2()
	fmt.Println("The function returen a and b ", a, b)
}

func saySomething() string {
	return "something"
}
func saySomething2() (string, string) {
	return "something", "else"
}
