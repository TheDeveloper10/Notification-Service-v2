package client

import "notification-service/internal/config"

type ISMS interface {
	SendSMS(title string, body string, to string) error
}

func InitSMSClient(conf *config.SMSConfig, real bool) ISMS {
	if real {
		return newRealSMSClientFromConfig(conf)
	} else {
		return newEmptySMSClientFromConfig()
	}
}
