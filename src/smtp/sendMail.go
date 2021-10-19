package smtp

import (
	"fmt"
	"net/smtp"
)

// Sends email to recipents in "to" array
// Takes smtpInfo, array of recipents, bytes.Bufffer,
func (s *SMTPInfo) SendEmail(index int) {

	auth := smtp.PlainAuth("", s.From, s.Password, s.Host)
	err := smtp.SendMail(s.Host+":"+s.Port, auth, s.From, s.To, []byte(s.Body))
	if err != nil {
		fmt.Println(err)
		return
	}
}
