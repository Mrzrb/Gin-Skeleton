package messages

import (
	"app/config"
	"net/smtp"
	"strconv"
)

type SmtpEmail struct {
	SMTPServer string
	FromName   string
	Password   string
	SMTPPort   int
}

func NewSmtpEmail(smtpServer, fromName, password string, smtpPort int) *SmtpEmail {
	return &SmtpEmail{
		SMTPServer: smtpServer,
		FromName:   fromName,
		Password:   password,
		SMTPPort:   smtpPort,
	}
}

func NewSmtpEmailNoreply() *SmtpEmail {
	cfg := config.Conf.EmailNoreply
	return NewSmtpEmail(cfg.EmailServer, cfg.EmailUser, cfg.EmailPassword, cfg.EmailPort)
}

// Send implements EmailProvider.
func (s *SmtpEmail) Send(email Email) error {
	// 组装邮件内容
	message := email.Content()

	// 连接到SMTP服务器
	auth := smtp.PlainAuth("", s.FromName, s.Password, s.SMTPServer)
	err := smtp.SendMail(s.SMTPServer+":"+strconv.Itoa(s.SMTPPort), auth, email.From, email.To, []byte(message))
	if err != nil {
		return err
	}
	return nil
}

var _ EmailProvider = (*SmtpEmail)(nil)
