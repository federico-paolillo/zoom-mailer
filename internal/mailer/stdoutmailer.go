package mailer

import (
	"fmt"
	"strings"
)

type StdoutMailer struct {
}

func (*StdoutMailer) Send(messageBody string, sendlist *Sendlist) error {

	recipients := strings.Join(sendlist.members, ";")

	fmt.Printf("Sending email to %s \n", recipients)
	fmt.Println(messageBody)

	return nil

}

func NewStdoutMailer() *StdoutMailer {
	return &StdoutMailer{}
}
