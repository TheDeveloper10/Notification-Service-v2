package client

import (
	"fmt"
	"net/smtp"
	"notification-service/internal/config"
	"notification-service/internal/util"
)

func newRealMailClientFromConfig(conf *config.MailConfig) *realMail {
	c := realMail{
		address:   fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		fromEmail: conf.FromEmail,
		auth:      smtp.PlainAuth("", conf.FromEmail, conf.FromPassword, conf.Host),
	}

	util.Logger.Info().Msg("Initialized Mail Client")
	return &c
}

type realMail struct {
	address   string
	fromEmail string
	auth      smtp.Auth
}

func (rm *realMail) Mail(subject string, message string, to []string) error {
	msg := fmt.Sprintf("Subject: %s\r\n\r\n%s", subject, message)
	return smtp.SendMail(rm.address, rm.auth, rm.fromEmail, to, []byte(msg))
}

func (rm *realMail) MailSingle(subject string, message string, to string) error {
	return rm.Mail(subject, message, []string{to})
}
