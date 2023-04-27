package dto

type Notification struct {
	AppID                 string            `json:"appId"`
	TemplateID            int               `json:"templateId"`
	Title                 string            `json:"title"`
	UniversalPlaceholders map[string]string `json:"placeholders"`

	Targets []NotificationTarget `json:"targets"`
}
