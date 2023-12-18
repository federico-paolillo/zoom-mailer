package mailer

type Mailer interface {
	Send(messageBody string, sendlist *Sendlist) error
}
