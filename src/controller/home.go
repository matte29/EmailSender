package controller

import (
	"fmt"
	"net/http"

	"github.com/matte29/EmailSender/src/smtp"
)

func Home(w http.ResponseWriter, r *http.Request) {

	var s *smtp.SMTPInfo

	if r.URL.Path != "/" {
		ErrorHandler(w, r, http.StatusNotFound)
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

		s.From = r.FormValue("from")
		fmt.Fprintf(w, "From = %s\n", s.From)

		s.Subject = r.FormValue("subject")
		fmt.Fprintf(w, "From = %s\n", s.Subject)

		s.Password = r.FormValue("password")
		fmt.Fprintf(w, "From = %s\n", s.Password)

		s.Host = r.FormValue("smtpHost")
		fmt.Fprintf(w, "From = %s\n", s.Host)

		s.Port = r.FormValue("smtpPort")
		fmt.Fprintf(w, "From = %s\n", s.Port)

		// TODO Read CSV, SetBody, SendMail

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}

}
