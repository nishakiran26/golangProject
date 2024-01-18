package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Fname     string `json:"first_name"`
	Lname     string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJson := `
	[
		{
			"first_name":"Clark",
			"last_name":"Kent",
			"hair_color":"black",
			"has_dog":true
		},
		{
			"first_name":"Bruce",
			"last_name":"Wayne",
			"hair_color":"black",
			"has_dog":false
		}
	]`

	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("err unmarshalling", err)

	}
	log.Printf("unmarshalled: %v", unmarshalled)

	//write json from struct

	var mySlice []Person
	var m1 Person
	m1.Fname = "Nisha"
	m1.Lname = "Rai"
	m1.HairColor = "Black"
	m1.HasDog = false

	mySlice = append(mySlice, m1)

	var m2 Person
	m2.Fname = "Prince"
	m2.Lname = "Sharma"
	m2.HairColor = "Red"
	m2.HasDog = false
	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "     ")
	if err != nil {
		log.Println("err marshal", err)
	}

	fmt.Println(string(newJson))

}
