package main

import (
	"net/http"
)

// home is homre page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home_page.html")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.html")
}
