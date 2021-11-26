package controller

import (
	"fmt"
	"io"
	"net/http"
	"os"

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
		fmt.Println("Get \"/\" was called.")
	case "POST":
		fmt.Println("Post \"/\" was called.")

		// Filename for CSV file
		var csvFileName string = "tmp/"

		// Filename for HTML file

		var htmlFileName string = "tmp/"

		if err := r.ParseMultipartForm(10 << 20); err != nil {
			fmt.Println("Error ParseMultipartForm: ", err)
			return
		}

		// Upload HTML File
		htmlFile, handler, err := r.FormFile("file")
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			return
		}

		defer htmlFile.Close()

		htmlFileName += handler.Filename

		dst, err := os.Create("tmp/" + handler.Filename)
		defer dst.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Copy the uploaded file to the created file on the filesystem
		if _, err := io.Copy(dst, htmlFile); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//--------------------------------------------------

		// Upload CSV File
		csvFile, csvHandler, err := r.FormFile("csv")
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			return
		}

		defer csvFile.Close()

		csvFileName += csvHandler.Filename

		copy, err := os.Create("tmp/" + csvHandler.Filename)
		defer copy.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Copy the uploaded file to the created file on the filesystem
		if _, err := io.Copy(copy, csvFile); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//--------------------------------------------------

		//! Something is wrong with getting these values
		s.From = r.FormValue("from")
		fmt.Printf("From = %s\n", s.From)

		s.Subject = r.FormValue("subject")
		fmt.Printf("Subject = %s\n", s.Subject)

		s.Password = r.FormValue("password")

		s.Host = r.FormValue("host")
		fmt.Printf("Host = %s\n", s.Host)

		s.Port = r.FormValue("port")
		fmt.Printf("Port = %s\n", s.Port)

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
		// TODO SetBody, SendMail
		for i, f := range s.To {
			s.SetBody(htmlFileName, i, f)

			s.SendEmail(i)
		}

	default:
		http.Redirect(w, r, "/404", 404)
	}

}
