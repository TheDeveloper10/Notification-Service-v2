package client

type ISMS interface {
	SendSMS(title string, body string, to string) error
}
