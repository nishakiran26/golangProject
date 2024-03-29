package main

import (
	"errors"
	"fmt"
	"net/http"
)

// *****************************************************
// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		n, err := fmt.Fprintf(w, "Hello world")
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
// 	})
// 	_ = http.ListenAndServe(":8080", nil)

// }
// *********************************************8Basic one
// main applicatipon function

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}

// home is homre page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "THis is home page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	//fmt.Fprintf(w, "This is about page")
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is a about page and 2+2 is %d", sum))
}

// add 2 integer and return sum
func addValues(x, y int) int {
	return x + y
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0")
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divide by %f is %f", 100.0, 0.0, f))
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot divide by 0")
		return 0, err
	}
	result := x / y
	return result, nil
}
