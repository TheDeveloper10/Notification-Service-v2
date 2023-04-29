package basic

import (
	"notification-service/internal/client"
	"notification-service/internal/dto"
	"notification-service/internal/util"
)

type NotificationSenderRepository struct {
}

func (nsr *NotificationSenderRepository) SendEmail(notification *dto.Notification) util.StatusCode {
	err := client.Mail.MailSingle(notification.Title, notification.Message, notification.ContactInfo)
	if err == nil {
		return util.StatusSuccess
	} else {
		return util.StatusError
	}
}

func (nsr *NotificationSenderRepository) SendSMS(notification *dto.Notification) util.StatusCode {
	err := client.SMS.SendSMS(notification.Title, notification.Message, notification.ContactInfo)
	if err == nil {
		return util.StatusSuccess
	} else {
		return util.StatusError
	}
}

func (nsr *NotificationSenderRepository) SendPush(notification *dto.Notification) util.StatusCode {
	err := client.Push.SendMessage(notification.Title, notification.Message, notification.ContactInfo)
	if err == nil {
		return util.StatusSuccess
	} else {
		return util.StatusError
	}
}
