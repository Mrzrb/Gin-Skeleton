package messages

type EmailProvider interface {
	Send(email Email) error
}

type Email struct {
	From    string
	Subject string
	Body    string
	To      []string
}

func NewEmail(from, subject, body string, to []string) Email {
	return Email{
		From:    from,
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
