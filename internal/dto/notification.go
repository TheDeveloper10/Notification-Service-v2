package dto

type Notification struct {
	ID uint64

	AppID       string
	TemplateID  uint64
	ContactInfo string
	Title       string
	Message     string

	SentTime uint32
}
