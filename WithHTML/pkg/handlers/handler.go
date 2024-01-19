package handlers

import (
	"net/http"

	"github.com/nishakiran26/go-learning/pkg/render"
)

// home is homre page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home_page.html")
	//renderTemplate(w, "home_page.html")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")
}
