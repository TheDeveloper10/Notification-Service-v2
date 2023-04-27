package repository

import (
	"notification-service/internal/dto"
	"notification-service/internal/repository/basic"
	"notification-service/internal/repository/mock"
	"notification-service/internal/util"
)

type INotificationSender interface {
	SendEmail(notification *dto.Notification) util.StatusCode
	SendSMS(notification *dto.Notification) util.StatusCode
	SendPush(notification *dto.Notification) util.StatusCode
}

func NewBasicNotificationSenderRepository() INotificationSender {
	return &basic.NotificationSenderRepository{}
}

func NewMockNotificationSenderRepository() INotificationSender {
	return &mock.NotificationSenderRepository{}
}
