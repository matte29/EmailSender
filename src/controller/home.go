package controller

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/matte29/EmailSender/src/csv"
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

		// Filename for CSV file
		var csvFileName string

		reader, err := r.MultipartReader()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//copy each part to destination.
		for {
			part, err := reader.NextPart()
			if err == io.EOF {
				break
			}

			//if part.FileName() is empty, skip this iteration.
			if part.FileName() == "" {
				continue
			}

			if strings.Contains(part.FileName(), ".csv") {
				csvFileName = part.FileName()
			}

			dst, err := os.Create("tmp/" + part.FileName())
			dst.Close()

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if _, err := io.Copy(dst, part); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		// Calls ReadCsvFile to get the emails from the CSV file
		s.To = csv.ReadCsvFile(csvFileName)

		if len(s.To) == 0 {
			fmt.Println("Error No Emails were recieved.")

			http.Error(w,
				"Error No Emails were recieved in the CSV File",
				http.StatusInternalServerError)

			return
		}
		// var wg sync.WaitGroup

		// var counter int = 0
		// for {
		// 	wg.Add(1)

		// }
		// TODO Read CSV, SetBody, SendMail

	default:
		http.Redirect(w, r, "/404", 404)
	}

}
