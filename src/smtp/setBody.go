package smtp

import (
	"bytes"
	"text/template"
)

// Adds template to a bytes.Buffer.  Takes the location of the HTML template,
// an interface{} that can be blank
//
//		Location: This is the location of the HTML file.
//		data: This is an interface, will be implemented further in the future,
//			  currently just set to the email address
//
//		TODO, Implement the "data: interface" based on user input

func (s *SMTPInfo) SetBody(location string, index int, data interface{}) error {

	t, err := template.ParseFiles(location)

	if err != nil {
		return err
	}

	s.Body += string(MIME)

	buffer := new(bytes.Buffer)
	if err = t.Execute(buffer, data); err != nil {
		return err
	}
	s.Body = buffer.String()

	s.Body = "To: " + s.To[index] + "\r\nSubject: " + s.Subject + "\r\n" + MIME + "\r\n" + s.Body
	return nil
}
