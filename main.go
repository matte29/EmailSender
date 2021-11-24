package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/matte29/EmailSender/src/controller"
)

func main() {

	http.HandleFunc("/", controller.Home)
	http.HandleFunc("/help", controller.Help)
	http.HandleFunc("/about", controller.About)
	http.HandleFunc("/404", controller.NotFoundPage)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
