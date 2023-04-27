package service

import (
	"notification-service/internal/dto"
	"notification-service/internal/util"
)

type Notification struct {
	templateSvc *Template
}

func (n *Notification) Send(notification *dto.Notification) (uint8, util.StatusCode) {
	template, status := n.templateSvc.GetTemplateByID(notification.TemplateID)
	if status != util.StatusSuccess {
		return 1, status
	}

	template.Body.Fill(notification.Placeholders)

	// for _, target := range notification.Targets {

	// }

	return 0, util.StatusSuccess
}
