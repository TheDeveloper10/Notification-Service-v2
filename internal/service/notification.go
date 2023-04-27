package service

import (
	"fmt"
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

	conHandler := concurrentHandler{
		wg:         sync.WaitGroup{},
		errorsChan: make(chan error),
	}

	conHandler.wg.Add(len(notificationReq.Placeholders))

	for index, target := range notificationReq.Targets {
		go n.handleTarget(
			&targetData{
				index:           index,
				target:          &target,
				notificationReq: notificationReq,
				template:        template,
			},
			&conHandler,
		)
	}

	conHandler.wg.Wait()

	return 0, util.StatusSuccess
}

func (n *Notification) handleTarget(tarData *targetData, conHandler *concurrentHandler) {
	defer conHandler.wg.Done()

	notificationTypes := []notificationType{
		{
			ContactInfo: tarData.target.Email,
			Body:        tarData.template.Body.Email,
			SendFunc:    n.notificationSenderRepo.SendEmail,
		},
		{
			ContactInfo: tarData.target.PhoneNumber,
			Body:        tarData.template.Body.SMS,
			SendFunc:    n.notificationSenderRepo.SendSMS,
		},
		{
			ContactInfo: tarData.target.FCMRegistrationToken,
			Body:        tarData.template.Body.Push,
			SendFunc:    n.notificationSenderRepo.SendPush,
		},
	}

	for _, nt := range notificationTypes {
		if nt.ContactInfo != nil && nt.Body != nil {
			bodyTemplate := util.TemplateString(*nt.Body)

			notification := dto.Notification{
				AppID:       tarData.notificationReq.AppID,
				TemplateID:  tarData.notificationReq.TemplateID,
				ContactInfo: *nt.ContactInfo,
				Title:       tarData.notificationReq.Title,
				Body:        bodyTemplate.Fill(tarData.target.Placeholders),
			}

			status := nt.SendFunc(&notification)
			if status != util.StatusSuccess {
				conHandler.errorsChan <- fmt.Errorf("Failed to send message for target %d for %s", tarData.index, *nt.ContactInfo)
				return
			}

			status = n.notificationRepo.SaveNotification(&notification)
			if status != util.StatusSuccess {
				conHandler.errorsChan <- fmt.Errorf("Failed to save sent message for target %d for %s", tarData.index, *nt.ContactInfo)
				return
			}
		}
	}
}

type notificationType struct {
	ContactInfo *string
	Body        *string
	SendFunc    func(*dto.Notification) util.StatusCode
}

type targetData struct {
	index           int
	notificationReq *dto.NotificationRequest
	target          *dto.NotificationTarget
	template        *dto.Template
}

type concurrentHandler struct {
	wg         sync.WaitGroup
	errorsChan chan error
}
