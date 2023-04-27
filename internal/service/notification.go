package service

import (
	"errors"
	"fmt"
	"notification-service/internal/dto"
	"notification-service/internal/repository"
	"notification-service/internal/util"
)

type Notification struct {
	templateSvc *Template

	notificationSenderRepo repository.INotificationSender
	notificationRepo       repository.INotification
}

func (n *Notification) Send(notificationReq *dto.NotificationRequest) ([]error, util.StatusCode) {
	template, status := n.templateSvc.GetTemplateByID(notificationReq.TemplateID)
	if status == util.StatusNotFound {
		return []error{errors.New("Template not found")}, util.StatusNotFound
	} else if status != util.StatusSuccess {
		return []error{errors.New("Failed to get template")}, util.StatusInternal
	}

	template.Body.Fill(notificationReq.Placeholders)

	se := newSyncErrors()

	se.wg.Add(len(notificationReq.Placeholders))

	for index, target := range notificationReq.Targets {
		td := &targetData{
			index:           index,
			target:          &target,
			notificationReq: notificationReq,
			template:        template,
		}

		go n.handleTarget(td, se)
	}

	se.wg.Wait()

	if se.errors == nil || len(se.errors) <= 0 {
		return nil, util.StatusSuccess
	}

	return se.errors, util.StatusError
}

func (n *Notification) handleTarget(tarData *targetData, se *syncErrors) {
	defer se.wg.Done()

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
		if nt.ContactInfo == nil || nt.Body == nil {
			continue
		}

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
			se.addError(fmt.Errorf("Failed to send message for target %d for %s", tarData.index, *nt.ContactInfo))
			return
		}

		status = n.notificationRepo.SaveNotification(&notification)
		if status != util.StatusSuccess {
			se.addError(fmt.Errorf("Failed to save sent message for target %d for %s", tarData.index, *nt.ContactInfo))
			return
		}
	}
}
