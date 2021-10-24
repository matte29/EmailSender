package main

import (
	"fmt"
	"log"
	"net/http"
)

func sender(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "templates/index.html")
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)

		from := r.FormValue("from")
		fmt.Fprintf(w, "From = %s\n", from)

		subject := r.FormValue("subject")
		fmt.Fprintf(w, "From = %s\n", subject)

		password := r.FormValue("password")
		fmt.Fprintf(w, "From = %s\n", password)

		fieldName := r.FormValue("fieldName")
		fmt.Fprintf(w, "From = %s\n", fieldName)

		smtpHost := r.FormValue("smtpHost")
		fmt.Fprintf(w, "From = %s\n", smtpHost)

		smtpPort := r.FormValue("smtpPort")
		fmt.Fprintf(w, "From = %s\n", smtpPort)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func help(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/help" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "templates/help.html")
	default:
		fmt.Fprintf(w, "404") // TODO Change to serve a 404 html file
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/about" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "templates/about.html")
	default:
		fmt.Fprintf(w, "404.") // TODO Change to serve a 404 html file
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "custom 404")
	}
}

func main() {
	http.HandleFunc("/", sender)
	http.HandleFunc("/help", help)
	http.HandleFunc("/about", about)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
