package dto

type Notification struct {
	AppID       string
	TemplateID  uint64
	ContactInfo string
	Title       string
	Body        string
}
