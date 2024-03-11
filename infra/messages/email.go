package messages

import "app/config"

type EmailProvider interface {
	Send(email Email) error
}

type Email struct {
	From    string
	Subject string
	Body    string
	To      []string
}

func NewEmail(subject, body string, to []string) Email {
	return Email{
		From:    config.Conf.EmailNoreply.EmailUser,
		To:      to,
		Subject: subject,
		Body:    body,
	}
}

func (email *Email) Content() string {
	message := "From: " + email.From + "\n" +
		"To: " + email.To[0] + "\n" +
		"Subject: " + email.Subject + "\n\n" +
		email.Body

	return message
}
