package smtp

import (
	"fmt"
	"net/smtp"
)

// Sends email to recipents in s.To array
// Takes smtpInfo, array of recipents, bytes.Bufffer,
func (s *SMTPInfo) SendEmail(index int) {

	var stringArray []string

	stringArray = append(stringArray, s.To[index])

	auth := smtp.PlainAuth("", s.From, s.Password, s.Host)
	err := smtp.SendMail(s.Host+":"+s.Port, auth, s.From, stringArray, []byte(s.Body))
	if err != nil {
		fmt.Println(err)
		return
	}
}
