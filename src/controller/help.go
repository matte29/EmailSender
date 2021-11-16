package controller

import (
	"fmt"
	"net/http"
)

func Help(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/help" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "templates/help.html")
	default:
		fmt.Fprintf(w, "404") // TODO Change to serve a 404 html file
	}
}
