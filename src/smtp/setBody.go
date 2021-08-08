package SMTP

import (
	"bytes"
	"text/template"
)

func setBody(location string, to string) bytes.Buffer {

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(mimeHeaders))

	t, _ := template.ParseFiles(location)

	t.Execute(&body, struct {
		Email string
	}{
		Email: to,
	})

	return body
}
