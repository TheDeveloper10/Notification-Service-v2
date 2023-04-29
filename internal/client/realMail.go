package client

type realMail struct {
}

func (rm *realMail) Mail(subject string, message string, to []string) error {
	return nil
}

func (rm *realMail) MailSingle(subject string, message string, to string) error {
	return nil
}
