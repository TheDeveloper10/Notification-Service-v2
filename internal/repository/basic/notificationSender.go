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
		util.Logger.Info().Msg("Sent an email notification")
		return util.StatusSuccess
	}
	return nsr.handleError(err)
}

func (nsr *NotificationSenderRepository) SendSMS(notification *dto.Notification) util.StatusCode {
	err := client.SMS.SendSMS(notification.Title, notification.Message, notification.ContactInfo)

	if err == nil {
		util.Logger.Info().Msg("Sent an SMS notification")
		return util.StatusSuccess
	}
	return nsr.handleError(err)
}

func (nsr *NotificationSenderRepository) SendPush(notification *dto.Notification) util.StatusCode {
	err := client.Push.SendMessage(notification.Title, notification.Message, notification.ContactInfo)

	if err == nil {
		util.Logger.Info().Msg("Sent a push notification")
		return util.StatusSuccess
	}
	return nsr.handleError(err)
}

func (nsr *NotificationSenderRepository) handleError(err error) util.StatusCode {
	util.Logger.Error().Msg(err.Error())
	return util.StatusError
}
