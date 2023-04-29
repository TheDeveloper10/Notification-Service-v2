package client

func newEmptySMSClientFromConfig() *emptySMS {
	return &emptySMS{}
}

type emptySMS struct {
}

func (es *emptySMS) SendSMS(title string, body string, to string) error {
	return nil
}
