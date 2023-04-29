package client

import "notification-service/internal/config"

func newEmptySMSClientFromConfig(conf *config.SMSConfig) *emptySMS {
	return &emptySMS{}
}

type emptySMS struct {
}

func (es *emptySMS) SendSMS(title string, body string, to string) error {
	return nil
}
