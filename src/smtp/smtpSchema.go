package smtp

type SMTPInfo struct {
	Host     string
	Port     string
	From     string
	Password string

	//* Email information for who recieves the Emails
	To      map[int]string
	Subject string
	Body    string
}

const (
	MIME string = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
)
