package service

import (
	"notification-service/internal/dto"
	"notification-service/internal/repository"
	"notification-service/internal/util"
	"sync"
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
	wg := sync.WaitGroup{}

	wg.Add(len(notificationReq.Targets))

	for _, target := range notificationReq.Targets {
		go n.handleTarget(notificationReq, &target, template, &wg)
	}

	wg.Wait()
	return 0, util.StatusSuccess
}

func (n *Notification) handleTarget(
	notificationReq *dto.NotificationRequest,
	target *dto.NotificationTarget,
	template *dto.Template,
	wg *sync.WaitGroup) {
	defer wg.Done()

	for _, notificationType := range []notificationType{
		{TargetInfo: target.Email, Body: template.Body.Email, SendFunc: n.notificationSenderRepo.SendEmail},
		{TargetInfo: target.PhoneNumber, Body: template.Body.SMS, SendFunc: n.notificationSenderRepo.SendSMS},
		{TargetInfo: target.FCMRegistrationToken, Body: template.Body.Push, SendFunc: n.notificationSenderRepo.SendPush},
	} {
		if notificationType.TargetInfo != nil && notificationType.Body != nil {
			bodyTemplate := util.TemplateString(*notificationType.Body)

			simpleNotification := dto.SimpleNotification{
				ContactInfo: *notificationType.TargetInfo,
				Title:       notificationReq.Title,
				Body:        bodyTemplate.Fill(target.Placeholders),
			}
			notificationType.SendFunc(&simpleNotification)
		}
	}
}

type notificationType struct {
	TargetInfo *string
	Body       *string
	SendFunc   func(*dto.SimpleNotification) util.StatusCode
}
