package smtp

import (
	"bytes"
	"fmt"
	"net/smtp"
)

// Sends email to recipents in "to" array
// Takes smtpInfo, array of recipents, bytes.Bufffer,
func SendEmail(s SMTPInfo, to []string, body bytes.Buffer) {

	auth := smtp.PlainAuth("", s.From, s.Password, s.Host)
	err := smtp.SendMail(s.Host+":"+s.Port, auth, s.From, to, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
}
