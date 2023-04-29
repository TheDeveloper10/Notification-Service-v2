package client

type IMail interface {
	Mail(subject string, message string, to []string) error
	MailSingle(subject string, message string, to string) error
}
