package main

import "log"

func main() {
	// for i := 0; i <= 5; i++ {
	// 	log.Println(i)
	// }

	animals := []string{"dog", "fish", "cow", "horse", "cat"}
	for i, animal := range animals {
		log.Println(i, animal)
	}
	for _, ani := range animals {
		log.Println(ani)
	}

	names := make(map[string]string)
	names["Nisha1"] = "Nisha"
	names["Kiran1"] = "Kiran"
	names["Rai1"] = "Rai"

	for nameType, name := range names {
		log.Println(nameType, name)
	}

	var firstLine = "Once upon a midnight dreary"

	for i, l := range firstLine {
		log.Println(i, ":", l) // this print numbers(byte) not letters- string is slice of byte
	}

	type User struct {
		Fname string
		Lname string
		Email string
		Age   int
	}

	var users []User
	users = append(users, User{"Nisha", "Rai", "xyz@gmail.com", 25})
	users = append(users, User{"Rupali", "Yadav", "abc@gmail.com", 21})
	users = append(users, User{"Kiran", "Sharma", "pqr@gmail.com", 24})
	users = append(users, User{"Ravi", "Singh", "jkl@gmail.com", 23})

	for _, l := range users {
		log.Println(l.Fname, l.Lname, l.Email, l.Age)
	}
}
