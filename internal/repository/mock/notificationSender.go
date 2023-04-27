package mock

import (
	"notification-service/internal/dto"
	"notification-service/internal/util"
)

type NotificationSenderRepository struct {
}

func (nsr *NotificationSenderRepository) SendEmail(notification *dto.Notification) util.StatusCode {
	return util.StatusSuccess
}

func (nsr *NotificationSenderRepository) SendSMS(notification *dto.Notification) util.StatusCode {
	return util.StatusSuccess
}

func (nsr *NotificationSenderRepository) SendPush(notification *dto.Notification) util.StatusCode {
	return util.StatusSuccess
}
