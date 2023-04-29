package client

import (
	"notification-service/internal/config"
)

type IMail interface {
	Mail(subject string, message string, to []string) error
	MailSingle(subject string, message string, to string) error
}

func InitMailClient(conf *config.MailConfig, empty bool) IMail {
	if empty {
		return newEmptyMailClientFromConfig(conf)
	} else {
		return newRealMailClientFromConfig(conf)
	}
}
