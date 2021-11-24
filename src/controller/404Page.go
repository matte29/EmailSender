package controller

import "net/http"

func NotFoundPage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "templates/404.html")

	default:
		http.Redirect(w, r, "/404", 404)

	}
}
