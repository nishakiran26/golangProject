package main

import (
	"log"
	"sort"
)

type User struct {
	Fname string
	Lname string
}

func main() {
	myMap := make(map[string]string)
	myMap["dog"] = "Samson"
	myMap["other-dog"] = "Cassie"
	myMap["dog"] = "fido"
	log.Println(myMap["dog"])
	log.Println(myMap["other-dog"])

	myMap2 := make(map[string]int)
	myMap2["First"] = 1
	myMap2["Second"] = 2
	log.Println(myMap2["First"], myMap2["Second"])
	myMap3 := make(map[string]User)
	me := User{
		Fname: "Nisha",
		Lname: "Rai",
	}
	myMap3["me"] = me
	log.Println(myMap3["me"].Fname)

	// var myNewVar float32
	// myNewVar=11.1

	// Slices

	var mySlice []string
	mySlice = append(mySlice, "Nisha")
	mySlice = append(mySlice, "Kiran")
	mySlice = append(mySlice, "Rupali")
	log.Println(mySlice)

	var mySlice2 []int
	mySlice2 = append(mySlice2, 2)
	mySlice2 = append(mySlice2, 1)
	mySlice2 = append(mySlice2, 3)
	log.Println(mySlice2)

	sort.Ints(mySlice2)
	log.Println(mySlice2)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	log.Println(numbers)
	log.Println(numbers[6:9])

	names := []string{"one", "seven", "fish", "cat"}
	log.Println(names)
}
