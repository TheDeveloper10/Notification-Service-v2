package controller

type TestRMQ struct {
}

func (ctrl *TestRMQ) Get(data []byte) (any, bool) {
	return nil, true
}
