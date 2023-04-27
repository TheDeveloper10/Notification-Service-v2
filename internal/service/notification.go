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

	if template.Body.Email != nil {

	}
	if template.Body.SMS != nil {

	}
	if template.Body.Push != nil {

	}

	return 0, util.StatusSuccess
}
