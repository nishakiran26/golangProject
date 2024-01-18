package main

import (
	"log"
	"time"
)

var s = "seven"

type User struct {
	FName  string
	LName  string
	Mobile string
	Age    int
	DOB    time.Time
}
type myStruct struct {
	FirstName string
}

func (m *myStruct) printFirstName() string {
	return m.FirstName
}
func main() {
	user := User{
		FName:  "Nisha",
		LName:  "Rai",
		Mobile: "8000000000",
	}
	log.Println(user.FName, user.LName, "DOB", user.DOB, "Mobile", user.Mobile)
	var myVar myStruct
	myVar.FirstName = "Rupali"

	myVar2 := myStruct{
		FirstName: "Ravi",
	}
	log.Println("myVar is set to ", myVar.printFirstName())
	log.Println("myVar2 is set to ", myVar2.printFirstName())
}

// func main() {
// 	var s2 = "six"
// 	log.Println("s is", s)
// 	log.Println("s2 is ", s2)
// 	saySomething("abc")

// }

// func saySomething(s3 string) (string, string) {
// 	log.Println("s from saysomething", s)
// 	return s, "World"
// }
