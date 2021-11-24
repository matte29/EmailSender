package controller

import (
	"net/http"
)

func About(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/about" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "templates/about.html")
	default:
		http.Redirect(w, r, "/404", 404) // TODO Change to serve a 404 html file
	}
}
