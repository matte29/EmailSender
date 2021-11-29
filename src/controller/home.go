package controller

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/matte29/EmailSender/src/csv"
	"github.com/matte29/EmailSender/src/smtp"
)

func post(w http.ResponseWriter, r *http.Request) {

	var s smtp.SMTPInfo

	// Filename for CSV file
	var csvFileName string = "tmp/"

	// Filename for HTML file

	var htmlFileName string = "tmp/"

	//! Maybe change this to larger size
	// Maximum upload of 10 MB files
	r.ParseMultipartForm(10 << 20)

	// Get data from the HTML Form
	s.From = r.FormValue("from")
	s.Subject = r.FormValue("subject")
	s.Password = r.FormValue("password")
	s.Host = r.FormValue("host")
	s.Port = r.FormValue("port")

	// Get and Upload HTML File
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)
	htmlFileName += handler.Filename

	// Create file
	dst, err := os.Create("tmp/" + handler.Filename)
	defer dst.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Copy the uploaded file to the created file on the filesystem
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//--------------------------------------------------

	// Get and Upload CSV File
	csvFile, handler, err := r.FormFile("csv")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	defer csvFile.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)
	csvFileName += handler.Filename

	// Create file
	dst2, err := os.Create("tmp/" + handler.Filename)
	defer dst2.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Copy the uploaded file to the created file on the filesystem
	if _, err := io.Copy(dst2, csvFile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//--------------------------------------------------

	fmt.Printf("Subject = %s\n", s.Subject)
	fmt.Printf("From = %s\n", s.From)
	fmt.Printf("Password = %s\n", s.Password)
	fmt.Printf("Host = %s\n", s.Host)
	fmt.Printf("Port = %s\n", s.Port)
	fmt.Println("Successfully Uploaded File")

	// Calls ReadCsvFile to get the emails from the CSV file
	s.To = csv.ReadCsvFile(csvFileName)

	if len(s.To) == 0 {
		fmt.Println("Error No Emails were recieved.")

		http.Error(w,
			"Error No Emails were recieved in the CSV File",
			http.StatusInternalServerError)

		return
	}

	// TODO: Change to the concurent way using waitgroup
	// Calls setBody and then sends the email to the email address
	for i, f := range s.To {
		s.SetBody(htmlFileName, i, f)

		s.SendEmail(i)
	}

	fmt.Println("Successfully Sent Emails")

}

func Home(w http.ResponseWriter, r *http.Request) {

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
		post(w, r)
	default:
		http.Redirect(w, r, "/404", 404)
	}

}
