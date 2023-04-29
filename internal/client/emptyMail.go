package client

import "notification-service/internal/config"

func newEmptyMailClientFromConfig(conf *config.MailConfig) *emptyMail {
	return &emptyMail{}
}

type emptyMail struct {
}

func (em *emptyMail) Mail(subject string, message string, to []string) error {
	return nil
}

func (em *emptyMail) MailSingle(subject string, message string, to string) error {
	return nil
}
