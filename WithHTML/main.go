package main

import (
	"fmt"
	"net/http"

	"github.com/nishakiran26/go-learning/pkg/handlers"
)

const portNumber = ":8081"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
