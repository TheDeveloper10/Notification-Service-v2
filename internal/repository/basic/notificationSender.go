package basic

import (
	"notification-service/internal/config"
	"notification-service/internal/dto"
	"notification-service/internal/util"
)

type NotificationSenderRepository struct {
}

func (nsr *NotificationSenderRepository) SendEmail(notification *dto.Notification) util.StatusCode {
	if !config.Master.Service.NotificationTypes.Has(config.NOTIFICATION_TYPE_MAIL) {
		return util.StatusSuccess
	}

	return util.StatusSuccess
}

func (nsr *NotificationSenderRepository) SendSMS(notification *dto.Notification) util.StatusCode {
	if !config.Master.Service.NotificationTypes.Has(config.NOTIFICATION_TYPE_SMS) {
		return util.StatusSuccess
	}

	return util.StatusSuccess
}

func (nsr *NotificationSenderRepository) SendPush(notification *dto.Notification) util.StatusCode {
	if !config.Master.Service.NotificationTypes.Has(config.NOTIFICATION_TYPE_PUSH) {
		return util.StatusSuccess
	}

	return util.StatusSuccess
}
