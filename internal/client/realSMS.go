package client

import "notification-service/internal/config"

func newRealSMSClientFromConfig(conf *config.SMSConfig) *realSMS {
	return &realSMS{}
}

type realSMS struct {
}

func (rs *realSMS) SendSMS(title string, body string, to string) error {
	return nil
}
