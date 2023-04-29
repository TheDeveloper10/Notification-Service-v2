package client

import "notification-service/internal/config"

type ISMS interface {
	SendSMS(title string, body string, to string) error
}

func InitSMSClient(conf *config.SMSConfig, empty bool) ISMS {
	if empty {
		return newEmptySMSClientFromConfig(conf)
	} else {
		return newRealSMSClientFromConfig(conf)
	}
}
