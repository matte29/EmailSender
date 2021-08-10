package sending

import (
	"bytes"
	"fmt"
	"net/smtp"
)

// Sends email to recipents in "to" array
// Takes smtpInfo, array of recipents, bytes.Bufffer,
func sendEmail(s smtpInfo, to []string, body bytes.Buffer) {

	auth := smtp.PlainAuth("", s.from, s.password, s.host)
	err := smtp.SendMail(s.host+":"+s.port, auth, s.from, to, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
}
