package service

import (
	"fmt"
	"notification-service/internal/dto"
	"notification-service/internal/repository"
	"notification-service/internal/util"
)

type Notification struct {
	templateSvc *Template

	notificationSenderRepo repository.INotificationSender
	notificationRepo       repository.INotification

	emailNotificationsStats dto.SuccessFailureStats
	smsNotificationsStats   dto.SuccessFailureStats
	pushNotificationsStats  dto.SuccessFailureStats
}

func (svc *Notification) SendNotifications(notificationReq *dto.NotificationRequest) ([]string, util.StatusCode) {
	template, status := svc.templateSvc.GetTemplateByID(notificationReq.TemplateID)
	if status == util.StatusNotFound {
		return []string{"Template not found"}, util.StatusNotFound
	} else if status != util.StatusSuccess {
		return []string{"Failed to get template"}, util.StatusInternal
	}

	template.Body.Fill(notificationReq.Placeholders)

	se := newSyncErrors()

	se.wg.Add(len(notificationReq.Targets))

	for index, target := range notificationReq.Targets {
		targetBuf := target
		td := &targetData{
			index:           index,
			target:          &targetBuf,
			notificationReq: notificationReq,
			template:        template,
		}

		go svc.handleTarget(td, se)
	}

	se.wg.Wait()

	if se.errors == nil || len(se.errors) <= 0 {
		return nil, util.StatusSuccess
	}

	return se.errors, util.StatusError
}

func (svc *Notification) handleTarget(tarData *targetData, se *syncErrors) {
	defer se.wg.Done()

	notificationTypes := []notificationType{
		{
			ContactInfo: tarData.target.Email,
			Message:     tarData.template.Body.Email,
			SendFunc:    svc.notificationSenderRepo.SendEmail,
			Stats:       &svc.emailNotificationsStats,
		},
		{
			ContactInfo: tarData.target.PhoneNumber,
			Message:     tarData.template.Body.SMS,
			SendFunc:    svc.notificationSenderRepo.SendSMS,
			Stats:       &svc.smsNotificationsStats,
		},
		{
			ContactInfo: tarData.target.FCMRegistrationToken,
			Message:     tarData.template.Body.Push,
			SendFunc:    svc.notificationSenderRepo.SendPush,
			Stats:       &svc.pushNotificationsStats,
		},
	}

	for _, nt := range notificationTypes {
		if nt.ContactInfo == nil || nt.Message == nil {
			continue
		}

		msgTemplate := util.TemplateString(*nt.Message)

		notification := dto.Notification{
			AppID:       tarData.notificationReq.AppID,
			TemplateID:  tarData.notificationReq.TemplateID,
			ContactInfo: *nt.ContactInfo,
			Title:       tarData.notificationReq.Title,
			Message:     msgTemplate.Fill(tarData.target.Placeholders),
		}

		status := nt.SendFunc(&notification)
		if status != util.StatusSuccess {
			se.addError(fmt.Sprintf("Failed to send message for target %d (%s)", tarData.index, *nt.ContactInfo))
			nt.Stats.AddFailures(1)
			continue
		}

		_, status = svc.notificationRepo.SaveNotification(&notification)
		if status != util.StatusSuccess {
			se.addError(fmt.Sprintf("Failed to save sent message for target %d (%s)", tarData.index, *nt.ContactInfo))
			nt.Stats.AddFailures(1)
			continue
		}

		nt.Stats.AddSuccesses(1)
	}
}

func (svc *Notification) GetBulkNotifications(filter *dto.NotificationBulkFilter) ([]dto.Notification, util.StatusCode) {
	return svc.notificationRepo.GetBulkNotifications(filter)
}

func (svc *Notification) GetEmailStats() *dto.SuccessFailureStats {
	return &svc.emailNotificationsStats
}

func (svc *Notification) GetSMSStats() *dto.SuccessFailureStats {
	return &svc.smsNotificationsStats
}

func (svc *Notification) GetPushStats() *dto.SuccessFailureStats {
	return &svc.pushNotificationsStats
}
