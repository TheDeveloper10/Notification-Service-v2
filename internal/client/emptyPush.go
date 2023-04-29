package client

func newEmptyPushClientFromConfig() *emptyPush {
	return &emptyPush{}
}

type emptyPush struct {
}

func (ep *emptyPush) SendMessage(title string, body string, to string) error {
	return nil
}
