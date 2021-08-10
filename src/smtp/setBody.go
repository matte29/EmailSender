package smtp

import (
	"bytes"
	"text/template"
)

// Adds template to a bytes.Buffer.  Takes the location of the HTML template,
// an interface{} that can be blank
func SetBody(location string, data interface{}) bytes.Buffer {

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(mimeHeaders))

	t, _ := template.ParseFiles(location)

	t.Execute(&body, data)
	return body
}
