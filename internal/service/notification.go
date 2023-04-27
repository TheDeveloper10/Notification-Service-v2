package service

import (
	"notification-service/internal/dto"
	"notification-service/internal/repository"
	"notification-service/internal/util"
)

type Notification struct {
	templateSvc *Template

	notificationSenderRepo repository.INotificationSender
	notificationRepo       repository.INotification
}

func (n *Notification) Send(notificationReq *dto.NotificationRequest) (uint8, util.StatusCode) {
	template, status := n.templateSvc.GetTemplateByID(notificationReq.TemplateID)
	if status != util.StatusSuccess {
		return 1, status
	}

	template.Body.Fill(notificationReq.Placeholders)

	for _, target := range notificationReq.Targets {
		if target.Email != nil && template.Body.Email != nil {
			go n.handleEmailTarget(notificationReq, &target, template)
		}

		if target.PhoneNumber != nil && template.Body.SMS != nil {

		}

		if target.FCMRegistrationToken != nil && template.Body.Push != nil {

		}
	}

	return 0, util.StatusSuccess
}

func (n *Notification) handleEmailTarget(
	notificationReq *dto.NotificationRequest,
	target *dto.NotificationTarget,
	template *dto.Template) {
	emailTemplate := util.TemplateString(*template.Body.Email)

	simpleNotification := dto.SimpleNotification{
		ContactInfo: *target.Email,
		Title:       notificationReq.Title,
		Body:        emailTemplate.Fill(target.Placeholders),
	}
	n.notificationSenderRepo.SendEmail(&simpleNotification)
}
