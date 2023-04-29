package client

type emptyMail struct {
}

func (em *emptyMail) Mail(subject string, message string, to []string) error {
	return nil
}

func (em *emptyMail) MailSingle(subject string, message string, to string) error {
	return nil
}
