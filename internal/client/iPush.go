package client

type IPush interface {
	SendMessage(title string, body string, to string) error
}
