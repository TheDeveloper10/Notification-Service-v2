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

	for _, target := range notificationReq.Targets {
		if target.Email != nil && template.Body.Email != nil {
			go n.handleEmailTarget(notificationReq, &target, template, &wg)
		}

		if target.PhoneNumber != nil && template.Body.SMS != nil {
			go n.handleSMSTarget(notificationReq, &target, template, &wg)
		}

		if target.FCMRegistrationToken != nil && template.Body.Push != nil {
			go n.handlePushTarget(notificationReq, &target, template, &wg)
		}
	}

	wg.Wait()
	return 0, util.StatusSuccess
}

func (n *Notification) handleEmailTarget(
	notificationReq *dto.NotificationRequest,
	target *dto.NotificationTarget,
	template *dto.Template,
	wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	bodyTemplate := util.TemplateString(*template.Body.Email)

	simpleNotification := dto.SimpleNotification{
		ContactInfo: *target.Email,
		Title:       notificationReq.Title,
		Body:        bodyTemplate.Fill(target.Placeholders),
	}
	n.notificationSenderRepo.SendEmail(&simpleNotification)
}

func (n *Notification) handleSMSTarget(
	notificationReq *dto.NotificationRequest,
	target *dto.NotificationTarget,
	template *dto.Template,
	wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	bodyTemplate := util.TemplateString(*template.Body.SMS)

	simpleNotification := dto.SimpleNotification{
		ContactInfo: *target.PhoneNumber,
		Title:       notificationReq.Title,
		Body:        bodyTemplate.Fill(target.Placeholders),
	}
	n.notificationSenderRepo.SendSMS(&simpleNotification)
}

func (n *Notification) handlePushTarget(
	notificationReq *dto.NotificationRequest,
	target *dto.NotificationTarget,
	template *dto.Template,
	wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	bodyTemplate := util.TemplateString(*template.Body.Push)

	simpleNotification := dto.SimpleNotification{
		ContactInfo: *target.FCMRegistrationToken,
		Title:       notificationReq.Title,
		Body:        bodyTemplate.Fill(target.Placeholders),
	}
	n.notificationSenderRepo.SendPush(&simpleNotification)
}
