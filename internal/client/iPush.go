package client

type IPush interface {
	SendMessage(title string, body string, to string) error
}

func InitPushClient(credentialsFile string, real bool) IPush {
	if real {
		return newRealPushClientFromConfig(credentialsFile)
	} else {
		return newEmptyPushClientFromConfig()
	}
}
