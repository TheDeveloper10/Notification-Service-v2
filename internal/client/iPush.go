package client

type IPush interface {
	SendMessage(title string, body string, to string) error
}

func InitPushClient(credentialsFile string, empty bool) IPush {
	if empty {
		return newEmptyPushClientFromConfig()
	} else {
		return newRealPushClientFromConfig(credentialsFile)
	}
}
