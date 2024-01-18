package main

import "log"

func main() {
	var isTrue bool
	isTrue = false

	if isTrue {
		log.Println("isTrue is", isTrue)
	} else {
		log.Println("is true is", isTrue)
	}

	cat := "cat2"
	if cat == "cat" {
		log.Println("cat is cat")
	} else {
		log.Println("cat is not cat")
	}

	myNum := 100
	isTrue2 := false

	if myNum > 99 && !isTrue2 {
		log.Println("My num is greater and isTrue2 is set to true")
	} else if myNum < 100 && isTrue2 {
		log.Println("1")
	} else if myNum == 101 || isTrue2 {
		log.Println("2")
	} else if myNum > 1000 && isTrue2 == false {
		log.Println("3")
	}

	//Switch case

	myVar := "Nisha"

	switch myVar {
	case "Nisha":
		log.Println("Nisha Is set to Nisha")

	case "Kiran":
		log.Println("Nisha is set to Kiran")
	case "Rupali":
		log.Println("Nisha is set to Rupali")
	default:
		log.Println("Nisha is not there")
	}
}
